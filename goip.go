package main

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var logger = zap.Must(zap.NewProduction())
var sugar = logger.Sugar()
var settings = LoadSettings()

func main() {
	s := &http.Server{
		Addr:           settings.Listen,
		Handler:        http.HandlerFunc(HttpGetRoot),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	sugar.Info("Starting server on ", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		sugar.Panic(err)
	}
}
