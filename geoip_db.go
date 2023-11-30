package main

import "github.com/oschwald/geoip2-golang"

func OpenGeoipDatabases() (geoip2.Reader, geoip2.Reader, geoip2.Reader) {
	// Open databases
	CityDB, err := geoip2.Open(settings.CityDB)
	if err != nil {
		logger.Fatal(err.Error())
	}
	CountryDB, err := geoip2.Open(settings.CountryDB)
	if err != nil {
		logger.Fatal(err.Error())
	}
	AsnDB, err := geoip2.Open(settings.AsnDB)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return *CityDB, *CountryDB, *AsnDB
}

var CityDB, CountryDB, AsnDB = OpenGeoipDatabases()
