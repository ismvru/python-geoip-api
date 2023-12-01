package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func LoadSettings() Settings {
	logger.Info("Loading settings...")
	config.WithOptions(config.ParseEnv, config.ParseDefault)
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("config.yml", "config.yaml")
	if err != nil {
		logger.Warn(err.Error())
	}
	settings := Settings{}
	err = config.BindStruct("", &settings)
	if err != nil {
		logger.Panic(err.Error())
	}
	return settings
}
