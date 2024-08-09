# Python GeoIP Responser

**Warning for GitHub Users**
This repository is a mirror of my private gitlab. You can open Issues here, but it is absolutely useless to make PR's.

- [Python GeoIP Responser](#python-geoip-responser)
  - [Configuration](#configuration)
    - [env variables (.env)](#env-variables-env)
  - [Endpoints](#endpoints)
    - [Endpoint `/api/v1/ip`](#endpoint-apiv1ip)
    - [Endpoint `/api/v1/ip/{ip}`](#endpoint-apiv1ipip)
  - [Respone format](#respone-format)
    - [Example](#example)
      - [Example - Correct response](#example---correct-response)
      - [Example - Invalid IP response](#example---invalid-ip-response)
    - [Schema](#schema)
      - [Schema - Correct response](#schema---correct-response)
      - [Schema - Invalid IP response](#schema---invalid-ip-response)
  - [Telegram bot configuration](#telegram-bot-configuration)
    - [Additional env variables (.env)](#additional-env-variables-env)
    - [Telegram conmmands and responses example](#telegram-conmmands-and-responses-example)

## Configuration

### env variables (.env)

```bash
# Bind to specific IP, or :: - IPv4 + IPv6, 0.0.0.0 - IPv4
HTTP_HOST=::
# Bind to port
HTTP_PORT=8000
# Path to GeoIP City database
GEOIP_CITY=geoip/GeoLite2-City.mmdb
# Path to GeoIP ASN database
GEOIP_ASN=geoip/GeoLite2-ASN.mmdb
```

## Endpoints

### Endpoint `/api/v1/ip`

Information about client IP, see [Example - Correct response](#example---correct-response)

### Endpoint `/api/v1/ip/{ip}`

Information about passed IP, in format `128.128.128.128` or `2001:2001::2001` see [Example - Correct response](#example---correct-response)

If passed IP is invalid - returns 422, see [Example - Invalid IP response](#example---invalid-ip-response)

## Respone format

### Example

#### Example - Correct response

[Schema - Correct response](#schema---correct-response)

```json
{
  "ip": "128.128.128.128",
  "country": "US",
  "country_name": "United States",
  "registered_country": "US",
  "registered_country_name": "United States",
  "continent": "NA",
  "continent_name": "North America",
  "asn": 11499,
  "org": "WHOI-WOODSHOLE",
  "properties": {
    "unspecified": false,
    "loopback": false,
    "private": false,
    "multicast": false,
    "link_local": false,
    "reserved": false
  },
  "tz": "America/New_York",
  "latitude": 41.5226,
  "longitude": -70.6662,
  "accuracy_radius": 5
}

```

#### Example - Invalid IP response

[Schema - Invalid IP response](#schema---invalid-ip-response)

```json
{
  "detail": [
    {
      "type": "ip_any_address",
      "loc": [
        "path",
        "ip"
      ],
      "msg": "value is not a valid IPv4 or IPv6 address",
      "input": "128.128.128.256"
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
    "ip": {
      "type": "string"
    },
    "country": {
      "type": "string"
    },
    "country_name": {
      "type": "string"
    },
    "registered_country": {
      "type": "string"
    },
    "registered_country_name": {
      "type": "string"
    },
    "continent": {
      "type": "string"
    },
    "continent_name": {
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
        "link_local": {
          "type": "boolean"
        },
        "reserved": {
          "type": "boolean"
        }
      },
      "required": [
        "unspecified",
        "loopback",
        "private",
        "multicast",
        "link_local",
        "reserved"
      ]
    },
    "tz": {
      "type": "string"
    },
    "latitude": {
      "type": "number"
    },
    "longitude": {
      "type": "number"
    },
    "accuracy_radius": {
      "type": "integer"
    }
  },
  "required": [
    "ip",
    "country",
    "country_name",
    "registered_country",
    "registered_country_name",
    "continent",
    "continent_name",
    "asn",
    "org",
    "properties",
    "tz",
    "latitude",
    "longitude",
    "accuracy_radius"
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
            "type": {
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
            "input": {
              "type": "string"
            }
          },
          "required": [
            "type",
            "loc",
            "msg",
            "input"
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

## Telegram bot configuration

### Additional env variables (.env)

```bash
# Telegram token from @BotFather
TELEGRAM_TOKEN=HereIsYourTelegramToken
# Whitelist
TELEGRAM_WHITELIST='["username1", "id1"]'
```

### Telegram conmmands and responses example

Command: /ip 128.128.128.128

Response:

```yaml
accuracy_radius: 5
asn: 11499
continent: NA
continent_name: North America
country: US
country_name: United States
ip: !!python/object/apply:ipaddress.IPv4Address
- 2155905152
latitude: 41.5694
longitude: -70.6152
org: WHOI-WOODSHOLE
properties:
  link_local: false
  loopback: false
  multicast: false
  private: false
  reserved: false
  unspecified: false
registered_country: US
registered_country_name: United States
tz: America/New_York
```
