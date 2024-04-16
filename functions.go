package main

import (
	"net"
)

func GetIPInfo(ip net.IP, c chan IpResponse) {
	CityRecord, err := CityDB.City(ip)
	if err != nil {
		logger.Panic(err.Error())
	}
	AsnRecord, err := AsnDB.ASN(ip)
	if err != nil {
		logger.Panic(err.Error())
	}
	c <- IpResponse{
		Ip:                     ip,
		Country:                CityRecord.Country.IsoCode,
		CountryName:            CityRecord.Country.Names["en"],
		RepresentedCountry:     CityRecord.RepresentedCountry.IsoCode,
		RepresentedCountryName: CityRecord.RepresentedCountry.Names["en"],
		RegisteredCountry:      CityRecord.RegisteredCountry.IsoCode,
		RegisteredCountryName:  CityRecord.RegisteredCountry.Names["en"],
		Continent:              CityRecord.Continent.Code,
		ContinentName:          CityRecord.Continent.Names["en"],
		City:                   CityRecord.City.Names["en"],
		ASN:                    int(AsnRecord.AutonomousSystemNumber),
		Org:                    AsnRecord.AutonomousSystemOrganization,
		Properties: IpProperties{
			IsUnspecified:             ip.IsUnspecified(),
			IsLoopback:                ip.IsLoopback(),
			IsPrivate:                 ip.IsPrivate(),
			IsMulticast:               ip.IsMulticast(),
			IsInterfaceLocalMulticast: ip.IsInterfaceLocalMulticast(),
			IsLinkLocalMulticast:      ip.IsLinkLocalMulticast(),
			IsLinkLocalUnicast:        ip.IsLinkLocalUnicast(),
			IsGlobalUnicast:           ip.IsGlobalUnicast(),
		},
		TZ:             CityRecord.Location.TimeZone,
		Latitude:       CityRecord.Location.Latitude,
		Longitude:      CityRecord.Location.Longitude,
		AccuracyRadius: CityRecord.Location.AccuracyRadius,
	}
}
