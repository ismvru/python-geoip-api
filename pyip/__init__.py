# Copyright (C) 2024 Mikhail Isaev <admin@ismv.ru>

import importlib.metadata
import logging
from ipaddress import ip_address

from fastapi import FastAPI, Request
from fastapi.exceptions import HTTPException
from fastapi.responses import RedirectResponse

# from fastapi_profiler import PyInstrumentProfilerMiddleware
from pydantic import IPvAnyAddress

from pyip.functions import GeoIPReader
from pyip.models import IpResponse
from pyip.settings import settings

logging.basicConfig(level=settings.log_level)

app = FastAPI(version=importlib.metadata.version("pyip"), title="pyip")

# app.add_middleware(
#     PyInstrumentProfilerMiddleware,
#     server_app=app,
#     profiler_output_type="html",
#     is_print_each_request=False,
#     open_in_browser=False,
#     html_file_name="pprof.html",
# )

reader = GeoIPReader(city_db=settings.geoip_city, asn_db=settings.geoip_asn)


@app.get("/favicon.ico", include_in_schema=False)
async def favicon():
    raise HTTPException(status_code=404, detail="Not found")


@app.get("/", include_in_schema=False)
async def get_root_redir(request: Request) -> RedirectResponse:
    return RedirectResponse(url="/api/v1/ip")


@app.get("/{ip}", include_in_schema=False)
async def get_ip_provided_redir(ip: IPvAnyAddress) -> RedirectResponse:
    return RedirectResponse(url=f"/api/v1/ip/{ip}")


@app.get("/api/v1/ip", response_model=IpResponse)
async def get_ip(request: Request) -> IpResponse:
    if request.client.host == "testclient":
        ip: IPvAnyAddress = IPvAnyAddress(ip_address("127.0.0.1"))
    else: # pragma: no cover
        ip: IPvAnyAddress = IPvAnyAddress(request.client.host)
    ip_response: IpResponse = await reader.get_ip_info(ip_address(ip))
    return ip_response


@app.get("/api/v1/ip/{ip}", response_model=IpResponse)
async def get_ip_provided(ip: IPvAnyAddress) -> IpResponse:
    with GeoIPReader(city_db=settings.geoip_city, asn_db=settings.geoip_asn) as reader:
        ip_response: IpResponse = await reader.get_ip_info(ip)
    return ip_response
