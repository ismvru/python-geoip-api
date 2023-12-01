package main

import (
	"errors"
	"net"
	"net/http"

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
