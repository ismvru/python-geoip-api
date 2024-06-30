from fastapi import FastAPI, Request
from .settings import settings
from .models import IpResponse
from pydantic import IPvAnyAddress
from ipaddress import ip_address
from .functions import GeoIP_Reader
import importlib.metadata

app = FastAPI(version=importlib.metadata.version("pyip"), title="pyip")
reader = GeoIP_Reader(city_db=settings.geoip_city, asn_db=settings.geoip_asn)


@app.get("/")
async def get_ip(request: Request) -> IpResponse:
    ip: IPvAnyAddress = ip_address(request.client.host)
    ip_response: IpResponse = await reader.get_ip_info(ip)
    return ip_response


@app.get("/{ip}")
async def get_ip_provided(ip: IPvAnyAddress) -> IpResponse:
    with GeoIP_Reader(city_db=settings.geoip_city, asn_db=settings.geoip_asn) as reader:
        ip_response: IpResponse = await reader.get_ip_info(ip)
    return ip_response
