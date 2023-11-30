package main

import (
	"errors"
	"net"
	"net/http"
)

func GetIPInfo(ip net.IP) (IpResponse, error) {
	// Init response
	resp := IpResponse{}
	resp.Ip = ip

	// Get info from city
	CityRecord, err := CityDB.City(ip)
	if err != nil {
		sugar.Panic(err)
		return resp, err
	}
	resp.City = CityRecord.City.Names["en"]

	// Get info from country
	CountryRecord, err := CountryDB.Country(ip)
	if err != nil {
		sugar.Panic(err)
		return resp, err
	}
	resp.Country = CountryRecord.Country.IsoCode
	resp.CountryName = CountryRecord.Country.Names["en"]

	// Get info from ASN
	AsnRecord, err := AsnDB.ASN(ip)
	if err != nil {
		sugar.Panic(err)
		return resp, err
	}
	resp.ASN = int(AsnRecord.AutonomousSystemNumber)
	resp.Org = AsnRecord.AutonomousSystemOrganization

	return resp, err

}

func ReadUserIP(r *http.Request) (net.IP, error) {
	IPAddress := net.ParseIP(r.Header.Get("X-Real-Ip"))
	if IPAddress == nil {
		IPAddress = net.ParseIP(r.Header.Get("X-Forwarded-For"))
	}
	if IPAddress == nil {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		IPAddress = net.ParseIP(host)
	}
	if IPAddress == nil {
		err := errors.New("can't get client ip")
		return IPAddress, err
	}
	return IPAddress, nil
}
