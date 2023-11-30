package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Settings struct {
	Listen    string `default:"${HTTP_LISTEN | :3333}"`
	CityDB    string `default:"${GEOIP_CITY | geoip/GeoLite2-City.mmdb}"`
	CountryDB string `default:"${GEOIP_COUNTRY | geoip/GeoLite2-Country.mmdb}"`
	AsnDB     string `default:"${GEOIP_ASN | geoip/GeoLite2-ASN.mmdb}"`
}

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
