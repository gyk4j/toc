FROM dockurr/samba

# VOLUME /storage

RUN apk update

WORKDIR /storage

ENV USER=user
ENV PASS=password

EXPOSE 137/udp
EXPOSE 138/udp
EXPOSE 139/tcp
EXPOSE 445/tcp