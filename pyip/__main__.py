import asyncio

import uvicorn


async def run_dev_server():
    config = uvicorn.Config("pyip:app", port=8000, log_level="info", reload=True)
    server = uvicorn.Server(config)
    await server.serve()


if __name__ == "__main__":
    asyncio.run(run_dev_server())
