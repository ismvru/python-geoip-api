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

#### Example - Invalid IP response

[Schema - Invalid IP response](#schema---invalid-ip-response)

```json
{
    "error": "invalid ip",
    "ip": "128.128.128.256",
    "client_ip": "128.128.128.128"
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
    "continent": {
      "type": "string"
    },
    "continent_name": {
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
    "continent",
    "continent_name",
    "city",
    "asn",
    "org",
    "properties"
  ]
}
```

#### Schema - Invalid IP response

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    },
    "ip": {
      "type": "string"
    },
    "client_ip": {
      "type": "string"
    }
  },
  "required": [
    "error",
    "ip",
    "client_ip"
  ]
}
```
