package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

var logger = zap.Must(zap.NewProduction())
var settings = LoadSettings()

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.GET("/", HttpGetRoot)
	logger.Sugar().Infof("Starting server on %s", settings.Listen)
	err := router.Run(settings.Listen)
	if err != nil {
		logger.Panic(err.Error())
	}
}
