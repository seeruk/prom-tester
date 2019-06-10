# Build steps.
FROM golang:1.12.5 AS builder

WORKDIR "/opt/go/src/github.com/seeruk/prom-tester"

ADD go.mod ./
ADD go.sum ./

RUN set -x \
    && go mod download

ADD . ./
RUN set -x \
    && CGO_ENABLED=0 go build -ldflags "-s -w" -o prom-tester ./cmd/ptsrv/main.go

# Artifact steps.
FROM alpine:latest

RUN set -x \
    && apk upgrade --no-cache \
    && apk --no-cache add ca-certificates tzdata \
    && addgroup prom-tester \
    && adduser -D -H -S -G prom-tester prom-tester \
    && mkdir -p /opt/prom-tester \
    && chown prom-tester: /opt/prom-tester

WORKDIR /opt/prom-tester

USER prom-tester

COPY --chown=prom-tester --from=builder \
    /opt/go/src/github.com/seeruk/prom-tester/prom-tester .

EXPOSE 8080

ENTRYPOINT [ "/opt/prom-tester/prom-tester" ]

