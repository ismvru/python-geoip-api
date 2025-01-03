# syntax=docker/dockerfile:1
# Copyright (C) 2024 Mikhail Isaev <admin@ismv.ru>


FROM python:3.12 AS builder

ENV PYTHONFAULTHANDLER=1 \
    PYTHONHASHSEED=random \
    PYTHONUNBUFFERED=1 \
    PIP_DEFAULT_TIMEOUT=100 \
    PIP_DISABLE_PIP_VERSION_CHECK=1 \
    PIP_NO_CACHE_DIR=1 \
    POETRY_VERSION=1.8.4

WORKDIR /code
COPY poetry.lock pyproject.toml /code/
RUN pip install "poetry==$POETRY_VERSION" && python -m venv /venv
COPY . /code/

# hadolint ignore=SC1091
RUN . /venv/bin/activate && poetry install --only main --no-root && poetry build

FROM python:3.12-slim
ENV WEB_CONCURRENCY=4
WORKDIR /app
COPY --from=builder /venv /venv
COPY --from=builder /code/dist /app/
COPY telegram.py /app/
USER nobody

# hadolint ignore=SC1091,DL3013
RUN . /venv/bin/activate && pip --no-cache-dir install -- *.whl
CMD ["/venv/bin/uvicorn", "pyip:app", "--host", "0.0.0.0", "--port", "8000", "--forwarded-allow-ips", "*"]
