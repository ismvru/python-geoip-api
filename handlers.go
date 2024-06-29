package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/gin-gonic/gin"
)

func HttpGetRoot(c *gin.Context) {
	UserIP := net.ParseIP(c.ClientIP())
	ch := make(chan IpResponse)
	go GetIPInfo(UserIP, ch)
	ipinfo := <-ch
	close(ch)
	c.IndentedJSON(http.StatusOK, ipinfo)
}

func HttpGetIp(c *gin.Context) {
	UserIP := net.ParseIP(c.Param("ip"))
	ClientIP := net.ParseIP(c.ClientIP())
	if UserIP == nil {
		err := errors.New("invalid ip")
		resp := InvalidIpResponse{err.Error(), c.Param("ip"), ClientIP}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}
	ch := make(chan IpResponse)
	go GetIPInfo(UserIP, ch)
	ipinfo := <-ch
	close(ch)
	c.IndentedJSON(http.StatusOK, ipinfo)
}

func HandleTelegramUpdates(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			logger.Sugar().Infof("[%s, %d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			WhiteList := strings.Split(settings.TelegramWhitelist, " ")
			if !slices.Contains(WhiteList, update.Message.From.UserName) && !slices.Contains(WhiteList, fmt.Sprint(update.Message.From.ID)) {
				msg.Text = fmt.Sprintf("User with name \"%s\" (ID \"%d\") not in whitelist!", update.Message.From.UserName, update.Message.From.ID)
			} else {
				switch update.Message.Command() {
				case "help":
					msg.Text = "I understand `/ip` and `/id` commands"
					msg.ParseMode = "markdown"
				case "id":
					msg.Text = fmt.Sprintf("UserName: `%s`\nID: `%d`", update.Message.From.UserName, update.Message.From.ID)
					msg.ParseMode = "markdown"
				case "ip":
					SplittedString := strings.Split(update.Message.Text, " ")
					if len(SplittedString) < 2 {
						msg.Text = "Please give me **valid** ip address"
						break
					}
					UserIP := net.ParseIP(SplittedString[1])
					if UserIP == nil {
						msg.Text = "Please give me **valid** ip address"
						break
					}
					ch := make(chan IpResponse)
					go GetIPInfo(UserIP, ch)
					ipinfo := <-ch
					YamlResponse, _ := yaml.Marshal(ipinfo)
					msg.Text = fmt.Sprintf("```yaml\n%s\n```", YamlResponse)
					msg.ParseMode = "markdown"
				default:
					msg.Text = "I don't know that command"
				}
			}
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
