FROM dockurr/samba

# RUN apk update

ARG S6_OVERLAY_VERSION=3.2.0.0
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

ARG S6_SERVICE=smb

RUN mkdir -p /etc/s6-overlay/s6-rc.d/${S6_SERVICE}
COPY s6-overlay/type /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/type
COPY s6-overlay/run /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/run
COPY s6-overlay/finish /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/finish

RUN mkdir -p touch /etc/s6-overlay/s6-rc.d/user/contents.d/
RUN touch /etc/s6-overlay/s6-rc.d/user/contents.d/${S6_SERVICE}
RUN mkdir -p /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/dependencies.d/
RUN touch /etc/s6-overlay/s6-rc.d/${S6_SERVICE}/dependencies.d/base

# Hard-coded to /storage in smb.conf, samba.sh and Dockerfile
WORKDIR /storage

ENV USER=user
ENV PASS=password

EXPOSE 137/udp
EXPOSE 138/udp
EXPOSE 139/tcp
EXPOSE 445/tcp

ENTRYPOINT ["/init"]
