FROM traefik:alpine

EXPOSE 8080

COPY config/traefik.toml /etc/traefik/traefik.toml