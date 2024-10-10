FROM toc-frontend AS frontend
FROM toc-backend AS backend
FROM alpine

RUN apk update && apk upgrade

ARG S6_OVERLAY_VERSION=3.2.0.0
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

WORKDIR /opt/toc/

COPY --from=backend /toc-backend/toc-backend.tar.gz /tmp/toc-backend.tar.gz
RUN tar -xzvf /tmp/toc-backend.tar.gz
COPY --from=frontend /app/toc-frontend.tar.gz /tmp/toc-frontend.tar.gz
RUN tar -C web -xzvf /tmp/toc-frontend.tar.gz

RUN mkdir -p /etc/s6-overlay/s6-rc.d/toc
COPY s6-overlay/type /etc/s6-overlay/s6-rc.d/toc/type
COPY s6-overlay/run /etc/s6-overlay/s6-rc.d/toc/run
COPY s6-overlay/finish /etc/s6-overlay/s6-rc.d/toc/finish

RUN mkdir -p touch /etc/s6-overlay/s6-rc.d/user/contents.d/
RUN touch /etc/s6-overlay/s6-rc.d/user/contents.d/toc
RUN mkdir -p /etc/s6-overlay/s6-rc.d/toc/dependencies.d/
RUN touch /etc/s6-overlay/s6-rc.d/toc/dependencies.d/base

ENTRYPOINT ["/init"]
