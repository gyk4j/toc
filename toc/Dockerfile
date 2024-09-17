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
ENTRYPOINT ["/usr/local/bin/toc"]
