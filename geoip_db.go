package main

import "github.com/oschwald/geoip2-golang"

func OpenGeoipDatabases() (geoip2.Reader, geoip2.Reader, geoip2.Reader) {
	// Open databases
	CityDB, err := geoip2.Open(settings.CityDB)
	if err != nil {
		sugar.Fatal(err)
	}
	CountryDB, err := geoip2.Open(settings.CountryDB)
	if err != nil {
		sugar.Fatal(err)
	}
	AsnDB, err := geoip2.Open(settings.AsnDB)
	if err != nil {
		sugar.Fatal(err)
	}
	return *CityDB, *CountryDB, *AsnDB
}

var CityDB, CountryDB, AsnDB = OpenGeoipDatabases()
