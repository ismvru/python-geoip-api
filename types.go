package main

import "net"

type InvalidIpResponse struct {
	Error    string `json:"error"`
	Ip       string `json:"ip"`
	ClientIP net.IP `json:"client_ip"`
}

type IpResponse struct {
	Ip                     net.IP       `json:"ip"`
	Country                string       `json:"country"`
	CountryName            string       `json:"country_name"`
	RepresentedCountry     string       `json:"represented_country"`
	RepresentedCountryName string       `json:"represented_country_name"`
	RegisteredCountry      string       `json:"registered_country"`
	RegisteredCountryName  string       `json:"registered_country_name"`
	Continent              string       `json:"continent"`
	ContinentName          string       `json:"continent_name"`
	City                   string       `json:"city"`
	ASN                    int          `json:"asn"`
	Org                    string       `json:"org"`
	Properties             IpProperties `json:"properties"`
	TZ                     string       `json:"tz"`
	Latitude               float64      `json:"latitude"`
	Longitude              float64      `json:"longitude"`
	AccuracyRadius         uint16       `json:"accuracy_radius"`
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
	Listen            string `default:"${HTTP_LISTEN | :3333}"`
	CityDB            string `default:"${GEOIP_CITY | geoip/GeoLite2-City.mmdb}"`
	AsnDB             string `default:"${GEOIP_ASN | geoip/GeoLite2-ASN.mmdb}"`
	TelegramToken     string `default:"${TELEGRAM_TOKEN}"`
	TelegramWhitelist string `default:"${TELEGRAM_WHITELIST}"`
}
