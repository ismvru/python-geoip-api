# Go GeoIP Responser

## Configuration

### config.yml or env variables

```yaml
listen: :3333  # env var: HTTP_LISTEN, IP and port to listen
citydb: geoip/GeoLite2-City.mmdb # env var: GEOIP_CITY, path to City database
countrydb: geoip/GeoLite2-Country.mmdb # env var: GEOIP_COUNTRY, path to Country database
asndb: geoip/GeoLite2-ASN.mmdb # env var: GEOIP_ASN, path to ASN database
```
