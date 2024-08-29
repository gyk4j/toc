FROM alpine:latest AS build

WORKDIR /opt/toc/

RUN apk update && apk upgrade && apk add go gcc musl musl-dev
RUN go version && /usr/bin/x86_64-alpine-linux-musl-gcc -v

COPY toc.go .
RUN CC=/usr/bin/x86_64-alpine-linux-musl-gcc \
  go build \
    --ldflags '-linkmode external -extldflags "-static"' \
    toc.go
# RUN ls -la && ./toc && sleep 10

FROM alpine:latest AS run

WORKDIR /opt/toc/

COPY --from=build /opt/toc/toc .
ENTRYPOINT ["./toc"]
