FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN apk update --no-cache && apk add --no-cache tzdata

RUN go build -o app ./cmd/main.go



FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow
WORKDIR /build

COPY --from=builder /build/app /build/app

COPY config/config.yaml /build/config/config.yaml
COPY migrations /build/migrations
CMD ["./app"]