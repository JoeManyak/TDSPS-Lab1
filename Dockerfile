FROM golang:alpine as build_base
RUN apk add build-base

WORKDIR /tmp/service

ADD . .

RUN go mod tidy
RUN go build -o ./out/app .

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=build_base /tmp/service/out/app /app/api

CMD ["/app/api"]