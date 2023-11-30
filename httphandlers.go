package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpGetRoot(c *gin.Context) {
	UserIP := net.ParseIP(c.ClientIP())
	ipinfo, err := GetIPInfo(UserIP)
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, ErrorResponse{"405 - Method not allowed!"})
	}
	c.IndentedJSON(http.StatusOK, ipinfo)
}
