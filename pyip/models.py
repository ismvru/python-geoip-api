from pydantic import BaseModel, IPvAnyAddress


class InvalidIpResponse(BaseModel):
    error: str
    ip: str
    client_ip: IPvAnyAddress


class IpProperties(BaseModel):
    unspecified: bool
    loopback: bool
    private: bool
    multicast: bool
    link_local: bool
    reserved: bool


class IpResponse(BaseModel):
    ip: IPvAnyAddress
    country: str | None = None
    country_name: str | None = None
    registered_country: str | None = None
    registered_country_name: str | None = None
    continent: str | None = None
    continent_name: str | None = None
    asn: int | None = None
    org: str | None = None
    properties: IpProperties
    tz: str | None = None
    latitude: float | None = None
    longitude: float | None = None
    accuracy_radius: int | None = None
