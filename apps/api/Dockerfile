FROM golang:1.22-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

FROM base AS development

RUN go install github.com/air-verse/air@v1.52.3

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

FROM base AS build

COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o /kin-api ./...
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /kin-api ./cmd/api

FROM alpine:3.20 AS runtime

WORKDIR /app

COPY --from=build /kin-api /usr/local/bin/kin-api

EXPOSE 8080

CMD ["kin-api"]