# Copyright (C) 2024 Mikhail Isaev <admin@ismv.ru>

from fastapi.testclient import TestClient

from pyip import app

client = TestClient(app=app, headers={"x-forwarded-for": "127.0.0.1"})


def test_read_ip():
    resp = client.get("/api/v1/ip")
    assert resp.status_code == 200
    assert resp.json() == {
        "ip": "127.0.0.1",
        "country": None,
        "country_name": None,
        "registered_country": None,
        "registered_country_name": None,
        "continent": None,
        "continent_name": None,
        "asn": None,
        "org": None,
        "properties": {
            "unspecified": False,
            "loopback": True,
            "private": True,
            "multicast": False,
            "link_local": False,
            "reserved": False,
        },
        "tz": None,
        "latitude": None,
        "longitude": None,
        "accuracy_radius": None,
    }


def test_read_ip_google():
    resp = client.get("/api/v1/ip/8.8.8.8")
    assert resp.status_code == 200
    assert resp.json() == {
        "ip": "8.8.8.8",
        "country": "US",
        "country_name": "United States",
        "registered_country": "US",
        "registered_country_name": "United States",
        "continent": "NA",
        "continent_name": "North America",
        "asn": 15169,
        "org": "GOOGLE",
        "properties": {
            "unspecified": False,
            "loopback": False,
            "private": False,
            "multicast": False,
            "link_local": False,
            "reserved": False,
        },
        "tz": "America/Chicago",
        "latitude": 37.751,
        "longitude": -97.822,
        "accuracy_radius": 1000,
    }
