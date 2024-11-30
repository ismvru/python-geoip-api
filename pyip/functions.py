# Copyright (C) 2024 Mikhail Isaev <admin@ismv.ru>

from ipaddress import IPv4Address, IPv6Address, ip_address
from typing import Self

import maxminddb
from aiocache import Cache, cached
from aiocache.serializers import PickleSerializer
from maxminddb.types import Record
from pydantic import FilePath, IPvAnyAddress

from pyip.models import IpProperties, IpResponse


class GeoIPReader:
    """GeoIP (mmdb2) database reader"""

    def __init__(self, city_db: FilePath, asn_db: FilePath) -> None:
        """Init GeoIP (mmdb2) database reader

        Args:
            city_db (FilePath): Path to city db
            asn_db (FilePath): Path to asn db
        """
        self.city_db_path = city_db
        self.asn_db_path = asn_db
        self.city_db: maxminddb.Reader = self._open(self.city_db_path)
        self.asn_db: maxminddb.Reader = self._open(self.asn_db_path)

    def _open(self, db: FilePath) -> maxminddb.Reader:
        """Open database and return reader

        Args:
            db (FilePath): database path

        Returns:
            maxminddb.Reader: reader
        """
        return maxminddb.open_database(db)

    def __del__(self) -> None:
        try:
            self.city_db.close()
            self.asn_db.close()
        except Exception:
            pass

    def __enter__(self) -> Self:
        return self

    def __exit__(self, type, value, traceback) -> None:
        try:
            self.city_db.close()
            self.asn_db.close()
        except Exception:
            pass
        if type is not None:
            print(
                f"Exception {type} occurred with value {value}, traceback: {traceback}"
            )
            return False
        return True

    @cached(ttl=10, cache=Cache.MEMORY, serializer=PickleSerializer())
    async def get_ip_info(self, ip: IPvAnyAddress) -> IpResponse:
        """Return IP address info

        Args:
            ip (IPvAnyAddress): IP address

        Returns:
            IpResponse: IP address info
        """
        city_info: Record | None = self.city_db.get(ip)
        asn_info: Record | None = self.asn_db.get(ip)
        ip_address_info: IPv4Address | IPv6Address = ip_address(ip)
        ip_properties: IpProperties = IpProperties.model_validate(
            {
                "unspecified": ip_address_info.is_unspecified,
                "loopback": ip_address_info.is_loopback,
                "private": ip_address_info.is_private,
                "multicast": ip_address_info.is_multicast,
                "link_local": ip_address_info.is_link_local,
                "reserved": ip_address_info.is_reserved,
            }
        )
        ip_response_template: dict = {
            "ip": ip,
            "properties": ip_properties,
        }

        if city_info is not None:
            try:
                ip_response_template["country"] = city_info["country"]["iso_code"]
                ip_response_template["country_name"] = city_info["country"]["names"][
                    "en"
                ]
            except KeyError:
                pass
            try:
                ip_response_template["registered_country"] = city_info[
                    "registered_country"
                ]["iso_code"]
                ip_response_template["registered_country_name"] = city_info[
                    "registered_country"
                ]["names"]["en"]
            except KeyError:
                pass
            try:
                ip_response_template["continent"] = city_info["continent"]["code"]
                ip_response_template["continent_name"] = city_info["continent"][
                    "names"
                ]["en"]
            except KeyError:
                pass
            try:
                ip_response_template["tz"] = city_info["location"]["time_zone"]
                ip_response_template["latitude"] = city_info["location"]["latitude"]
                ip_response_template["longitude"] = city_info["location"]["longitude"]
                ip_response_template["accuracy_radius"] = city_info["location"][
                    "accuracy_radius"
                ]
            except KeyError:
                pass

        if asn_info is not None:
            ip_response_template["asn"] = asn_info["autonomous_system_number"]
            ip_response_template["org"] = asn_info["autonomous_system_organization"]

        ip_response: IpResponse = IpResponse.model_validate(ip_response_template)

        return ip_response
