#!/usr/bin/env python3
# Copyright (C) 2024 Mikhail Isaev <admin@ismv.ru>

import asyncio
import logging

import yaml
from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.filters import CommandStart
from aiogram.filters.command import Command
from aiogram.types import Message

from pyip.functions import GeoIPReader
from pyip.models import IpResponse
from pyip.settings import settings

dp = Dispatcher()
reader = GeoIPReader(city_db=settings.geoip_city, asn_db=settings.geoip_asn)

GTFO_MESSAGE: str = "Not authorized"


@dp.message(CommandStart())
async def command_start_handler(message: Message) -> None:
    """
    This handler receives messages with `/start` command
    """
    logging.info(f"New msg from {message.from_user.id}: {message.text}")
    if (
        message.from_user.username not in settings.telegram_whitelist
        and message.from_user.id not in settings.telegram_whitelist
    ):
        await message.answer(GTFO_MESSAGE)
        return
    await message.answer("Hello!")


@dp.message(Command("id"))
async def command_id_handler(message: Message) -> None:
    logging.info(f"New msg from {message.from_user.id}: {message.text}")
    if (
        message.from_user.username not in settings.telegram_whitelist
        and message.from_user.id not in settings.telegram_whitelist
    ):
        await message.answer(GTFO_MESSAGE)
        return
    await message.answer(
        f"Name: {message.from_user.username}, ID: {message.from_user.id}"
    )


@dp.message(Command("ip"))
async def command_ip_handler(message: Message) -> None:
    logging.info(f"New msg from {message.from_user.id}: {message.text}")
    if (
        message.from_user.username not in settings.telegram_whitelist
        and message.from_user.id not in settings.telegram_whitelist
    ):
        await message.answer(GTFO_MESSAGE)
        return
    splitted_text: list[str] = message.text.split(" ")
    if len(splitted_text) != 2:
        await message.answer("Please exactly one IP")
        return
    ip_info: IpResponse = await reader.get_ip_info(splitted_text[1])
    await message.answer(f"```yaml\n{yaml.dump(ip_info.model_dump())}\n```")


async def main() -> None:
    logging.info("Init telegram bot")
    bot = Bot(
        token=settings.telegram_token,
        default=DefaultBotProperties(parse_mode=ParseMode.MARKDOWN),
    )
    await dp.start_polling(bot)


if __name__ == "__main__":
    logging.basicConfig(level=settings.log_level)
    asyncio.run(main())
