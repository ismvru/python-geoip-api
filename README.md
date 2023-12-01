# Go GeoIP Responser

- [Go GeoIP Responser](#go-geoip-responser)
  - [Configuration](#configuration)
    - [config.yml or env variables](#configyml-or-env-variables)
  - [Respone format](#respone-format)
    - [Example](#example)
    - [Schema](#schema)

## Configuration

### config.yml or env variables

```yaml
listen: :3333  # env var: HTTP_LISTEN, IP and port to listen
citydb: geoip/GeoLite2-City.mmdb # env var: GEOIP_CITY, path to City database
countrydb: geoip/GeoLite2-Country.mmdb # env var: GEOIP_COUNTRY, path to Country database
asndb: geoip/GeoLite2-ASN.mmdb # env var: GEOIP_ASN, path to ASN database
```

## Respone format

### Example

```json
{
    "ip": "128.128.128.128",
    "country": "US",
    "country_name": "United States",
    "city": "Buzzards Bay",
    "asn": 11499,
    "org": "WHOI-WOODSHOLE",
    "properties": {
        "unspecified": false,
        "loopback": false,
        "private": false,
        "multicast": false,
        "interface_local_multicast": false,
        "link_local_multicast": false,
        "link_local_unicast": false,
        "global_unicast": true
    }
}
```

### Schema

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "ip": {
      "type": "string"
    },
    "country": {
      "type": "string"
    },
    "country_name": {
      "type": "string"
    },
    "city": {
      "type": "string"
    },
    "asn": {
      "type": "integer"
    },
    "org": {
      "type": "string"
    },
    "properties": {
      "type": "object",
      "properties": {
        "unspecified": {
          "type": "boolean"
        },
        "loopback": {
          "type": "boolean"
        },
        "private": {
          "type": "boolean"
        },
        "multicast": {
          "type": "boolean"
        },
        "interface_local_multicast": {
          "type": "boolean"
        },
        "link_local_multicast": {
          "type": "boolean"
        },
        "link_local_unicast": {
          "type": "boolean"
        },
        "global_unicast": {
          "type": "boolean"
        }
      },
      "additionalProperties": true,
      "required": [
        "unspecified",
        "loopback",
        "private",
        "multicast",
        "interface_local_multicast",
        "link_local_multicast",
        "link_local_unicast",
        "global_unicast"
      ]
    }
  },
  "additionalProperties": true,
  "required": [
    "ip",
    "country",
    "country_name",
    "city",
    "asn",
    "org",
    "properties"
  ]
}
```
