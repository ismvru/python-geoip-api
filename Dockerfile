# syntax=docker/dockerfile:1

FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /goip
CMD ["/goip"]

FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /
COPY --from=builder /goip /goip
USER nonroot:nonroot
ENV http_listen=:3333
ENTRYPOINT ["/goip"]
