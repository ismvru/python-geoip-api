package main

import (
	"net"
)

func GetIPInfo(ip net.IP, c chan IpResponse) (IpResponse, error) {
	// Init response
	resp := IpResponse{}
	resp.Ip = ip
	resp.Properties.IsUnspecified = ip.IsUnspecified()
	resp.Properties.IsLoopback = ip.IsLoopback()
	resp.Properties.IsPrivate = ip.IsPrivate()
	resp.Properties.IsMulticast = ip.IsMulticast()
	resp.Properties.IsInterfaceLocalMulticast = ip.IsInterfaceLocalMulticast()
	resp.Properties.IsLinkLocalMulticast = ip.IsLinkLocalMulticast()
	resp.Properties.IsLinkLocalUnicast = ip.IsLinkLocalUnicast()
	resp.Properties.IsGlobalUnicast = ip.IsGlobalUnicast()

	// Get info from city
	CityRecord, err := CityDB.City(ip)
	if err != nil {
		logger.Panic(err.Error())
		return resp, err
	}
	resp.City = CityRecord.City.Names["en"]
	resp.Country = CityRecord.Country.IsoCode
	resp.CountryName = CityRecord.Country.Names["en"]
	resp.Continent = CityRecord.Continent.Code
	resp.ContinentName = CityRecord.Continent.Names["en"]

	// Get info from ASN
	AsnRecord, err := AsnDB.ASN(ip)
	if err != nil {
		logger.Panic(err.Error())
		return resp, err
	}
	resp.ASN = int(AsnRecord.AutonomousSystemNumber)
	resp.Org = AsnRecord.AutonomousSystemOrganization

	c <- resp
	return resp, err

}
