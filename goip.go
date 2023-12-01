package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var logger = zap.Must(zap.NewProduction())
var settings = LoadSettings()
var CityDB, AsnDB = OpenGeoipDatabases()

func main() {
	// Watch signals and close
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigs
		CloseGeoipDatabases()
		os.Exit(0)
	}()

	// Gin settings
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery())
	router.GET("/", HttpGetRoot)

	// Start
	logger.Sugar().Infof("Starting server on %s", settings.Listen)
	err := router.Run(settings.Listen)
	if err != nil {
		logger.Panic(err.Error())
	}
	defer CloseGeoipDatabases()
}
