FROM alpine:latest AS build

RUN apk update && apk upgrade && apk add go gcc musl musl-dev
RUN go version && /usr/bin/x86_64-alpine-linux-musl-gcc -v

WORKDIR /usr/local/src

COPY server.bak/toc.go .
RUN CC=/usr/bin/x86_64-alpine-linux-musl-gcc \
  go build -o /usr/local/bin/toc toc.go
# RUN ls -la && ./toc && sleep 10

FROM alpine:latest AS run

RUN apk update && apk upgrade && apk add iputils docker-cli docker-cli-compose

ARG S6_OVERLAY_VERSION=3.2.0.0
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

VOLUME /opt/toc/

WORKDIR /opt/toc/

#CMD ping -c 1 db ; \
#    ping -c 1 file ; \
#    ping -c 1 mail ; \
#    ping -c 1 web ; \
#CMD docker compose exec db cat /etc/hostname ; \
#    docker compose exec file cat /etc/hostname ; \
#    docker compose exec mail cat /etc/hostname ; \
#    docker compose exec web cat /etc/hostname ; \
#    sleep 120

COPY --from=build /usr/local/bin/toc /usr/local/bin/toc
# ENTRYPOINT ["/usr/local/bin/toc"]

# CMD ["/usr/local/bin/toc"]

RUN mkdir -p /etc/s6-overlay/s6-rc.d/toc
COPY s6-overlay/type /etc/s6-overlay/s6-rc.d/toc/type
COPY s6-overlay/run /etc/s6-overlay/s6-rc.d/toc/run
COPY s6-overlay/finish /etc/s6-overlay/s6-rc.d/toc/finish

RUN mkdir -p touch /etc/s6-overlay/s6-rc.d/user/contents.d/
RUN touch /etc/s6-overlay/s6-rc.d/user/contents.d/toc
RUN mkdir -p /etc/s6-overlay/s6-rc.d/toc/dependencies.d/
RUN touch /etc/s6-overlay/s6-rc.d/toc/dependencies.d/base

ENTRYPOINT ["/init"]
