# Go GeoIP Responser

- [Go GeoIP Responser](#go-geoip-responser)
  - [Configuration](#configuration)
    - [config.yml or env variables](#configyml-or-env-variables)
  - [Endpoints](#endpoints)
    - [Endpoint `/`](#endpoint-)
    - [Endpoint `/:ip`](#endpoint-ip)
  - [Respone format](#respone-format)
    - [Example](#example)
      - [Example - Correct response](#example---correct-response)
      - [Example - Invalid IP response](#example---invalid-ip-response)
    - [Schema](#schema)
      - [Schema - Correct response](#schema---correct-response)
      - [Schema - Invalid IP response](#schema---invalid-ip-response)

## Configuration

### config.yml or env variables

```yaml
listen: :3333  # env var: HTTP_LISTEN, IP and port to listen
citydb: geoip/GeoLite2-City.mmdb # env var: GEOIP_CITY, path to City database
asndb: geoip/GeoLite2-ASN.mmdb # env var: GEOIP_ASN, path to ASN database
```

## Endpoints

### Endpoint `/`

Information about client IP, see [Example - Correct response](#example---correct-response)

### Endpoint `/:ip`

Information about passed IP, in format `/128.128.128.128` or `2001:2001::2001` see [Example - Correct response](#example---correct-response)

If passed IP is invalid - returns 400, see [Example - Invalid IP response](#example---invalid-ip-response)

## Respone format

### Example

#### Example - Correct response

[Schema - Correct response](#schema---correct-response)

```json
{
    "accuracy_radius": 5,
    "asn": 11499,
    "continent": "NA",
    "continent_name": "North America",
    "country": "US",
    "country_name": "United States",
    "ip": "128.128.128.128",
    "latitude": 41.5694,
    "longitude": -70.6152,
    "org": "WHOI-WOODSHOLE",
    "properties": {
        "link_local": false,
        "loopback": false,
        "multicast": false,
        "private": false,
        "reserved": false,
        "unspecified": false
    },
    "registered_country": "US",
    "registered_country_name": "United States",
    "tz": "America/New_York"
}
```

#### Example - Invalid IP response

[Schema - Invalid IP response](#schema---invalid-ip-response)

```json
{
    "detail": [
        {
            "input": "128.128.128.256",
            "loc": [
                "path",
                "ip"
            ],
            "msg": "value is not a valid IPv4 or IPv6 address",
            "type": "ip_any_address"
        }
    ]
}
```

### Schema

#### Schema - Correct response

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "accuracy_radius": {
      "type": "integer"
    },
    "asn": {
      "type": "integer"
    },
    "continent": {
      "type": "string"
    },
    "continent_name": {
      "type": "string"
    },
    "country": {
      "type": "string"
    },
    "country_name": {
      "type": "string"
    },
    "ip": {
      "type": "string"
    },
    "latitude": {
      "type": "number"
    },
    "longitude": {
      "type": "number"
    },
    "org": {
      "type": "string"
    },
    "properties": {
      "type": "object",
      "properties": {
        "link_local": {
          "type": "boolean"
        },
        "loopback": {
          "type": "boolean"
        },
        "multicast": {
          "type": "boolean"
        },
        "private": {
          "type": "boolean"
        },
        "reserved": {
          "type": "boolean"
        },
        "unspecified": {
          "type": "boolean"
        }
      },
      "required": [
        "link_local",
        "loopback",
        "multicast",
        "private",
        "reserved",
        "unspecified"
      ]
    },
    "registered_country": {
      "type": "string"
    },
    "registered_country_name": {
      "type": "string"
    },
    "tz": {
      "type": "string"
    }
  },
  "required": [
    "accuracy_radius",
    "asn",
    "continent",
    "continent_name",
    "country",
    "country_name",
    "ip",
    "latitude",
    "longitude",
    "org",
    "properties",
    "registered_country",
    "registered_country_name",
    "tz"
  ]
}
```

#### Schema - Invalid IP response

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "detail": {
      "type": "array",
      "items": [
        {
          "type": "object",
          "properties": {
            "input": {
              "type": "string"
            },
            "loc": {
              "type": "array",
              "items": [
                {
                  "type": "string"
                },
                {
                  "type": "string"
                }
              ]
            },
            "msg": {
              "type": "string"
            },
            "type": {
              "type": "string"
            }
          },
          "required": [
            "input",
            "loc",
            "msg",
            "type"
          ]
        }
      ]
    }
  },
  "required": [
    "detail"
  ]
}
```
