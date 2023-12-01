package main

import (
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
