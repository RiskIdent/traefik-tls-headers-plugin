ARG TRAEFIK_VERSION=v3.0.0
ARG BASE_IMAGE=docker.io/traefik:${TRAEFIK_VERSION}
FROM ${BASE_IMAGE}

COPY testconfig/traefik.yml /etc/traefik/traefik.yml
COPY testconfig/dynamic.yml /etc/traefik/dynamic.yml

COPY . plugins-local/src/github.com/RiskIdent/traefik-tls-headers-plugin
