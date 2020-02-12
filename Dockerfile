FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git \
 && go build -o main

FROM alpine

WORKDIR /app

COPY --from=build /go/app/main .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/main

CMD ["./main"]