FROM golang:alpine AS base

ADD / /kt
WORKDIR /kt
RUN apk add git
RUN go build -o kt *.go

# ====
FROM alpine

RUN apk add ca-certificates

COPY --from=base /kt/kt /kt

ENTRYPOINT /kt
