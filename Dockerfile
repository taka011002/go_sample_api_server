FROM golang:1.12-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git \
 && go build cmd/api/main.go

FROM alpine

WORKDIR /app

COPY --from=build /go/app/main .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/main

CMD ["./app"]