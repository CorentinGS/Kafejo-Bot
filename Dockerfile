FROM golang:1.22-alpine as builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update && apk add tzdata && apk add upx

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go get -d -v
RUN go build -ldflags="-s -w" -o /app/kafejobot .
RUN upx /app/kafejobot

FROM alpine:3.17

RUN apk update && apk add ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Paris /usr/share/zoneinfo/Europe/Paris
ENV TZ Europe/Paris

WORKDIR /app

COPY --from=builder /app/kafejobot /app/kafejobot
COPY --from=builder /build/.env /app/.env

EXPOSE 1813

CMD ["/app/kafejobot"]