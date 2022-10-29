FROM golang:alpine as build_base

WORKDIR /tmp/service

ADD . .

RUN go build -o ./out/app .

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=build_base /tmp/service/out/app /app/api

CMD ["/app/api"]