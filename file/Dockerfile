FROM dockurr/samba

# VOLUME /media/storage

RUN apk update

WORKDIR /media/storage

ENV USER=user
ENV PASS=password

EXPOSE 137/udp
EXPOSE 138/udp
EXPOSE 139/tcp
EXPOSE 445/tcp