from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import IPvAnyAddress, FilePath
from ipaddress import ip_address


class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_file=".env", env_file_encoding="utf-8")
    http_host: IPvAnyAddress = ip_address("::")
    http_port: int = 8000
    geoip_city: FilePath = "/geoip/GeoLite2-City.mmdb"
    geoip_asn: FilePath = "/geoip/GeoLite2-ASN.mmdb"
    telegram_token: str | None = None
    telegram_whitelist: list[str | int] | None = None


settings = Settings()
