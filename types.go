package main

import "net"

type IpResponse struct {
	Ip          net.IP `json:"ip"`
	Country     string `json:"country"`
	CountryName string `json:"country_name"`
	City        string `json:"city"`
	ASN         int    `json:"asn"`
	Org         string `json:"org"`
}
