FROM yobasystems/alpine-mariadb:latest

ARG S6_OVERLAY_VERSION=3.2.0.0
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

ARG S6_SERVICE=mysql

RUN mkdir -p /etc/s6-overlay/s6-rc.d/${S6_SERVICE}
COPY s6-overlay/type /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/type
COPY s6-overlay/run /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/run
COPY s6-overlay/finish /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/finish

RUN mkdir -p touch /etc/s6-overlay/s6-rc.d/user/contents.d/
RUN touch /etc/s6-overlay/s6-rc.d/user/contents.d/${S6_SERVICE}
RUN mkdir -p /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/dependencies.d/
RUN touch /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/dependencies.d/base

ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=db
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes
ENV MYSQL_RANDOM_ROOT_PASSWORD=
ENV MYSQL_ONETIME_PASSWORD=

ENTRYPOINT ["/init"]

