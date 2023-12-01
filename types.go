package main

import "net"

type IpResponse struct {
	Ip          net.IP       `json:"ip"`
	Country     string       `json:"country"`
	CountryName string       `json:"country_name"`
	City        string       `json:"city"`
	ASN         int          `json:"asn"`
	Org         string       `json:"org"`
	Properties  IpProperties `json:"properties"`
}

type IpProperties struct {
	IsUnspecified             bool `json:"unspecified"`
	IsLoopback                bool `json:"loopback"`
	IsPrivate                 bool `json:"private"`
	IsMulticast               bool `json:"multicast"`
	IsInterfaceLocalMulticast bool `json:"interface_local_multicast"`
	IsLinkLocalMulticast      bool `json:"link_local_multicast"`
	IsLinkLocalUnicast        bool `json:"link_local_unicast"`
	IsGlobalUnicast           bool `json:"global_unicast"`
}

type Settings struct {
	Listen    string `default:"${HTTP_LISTEN | :3333}"`
	CityDB    string `default:"${GEOIP_CITY | geoip/GeoLite2-City.mmdb}"`
	CountryDB string `default:"${GEOIP_COUNTRY | geoip/GeoLite2-Country.mmdb}"`
	AsnDB     string `default:"${GEOIP_ASN | geoip/GeoLite2-ASN.mmdb}"`
}
