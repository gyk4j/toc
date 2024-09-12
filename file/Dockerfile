FROM dockurr/samba

# RUN apk update

# Hard-coded to /storage in smb.conf, samba.sh and Dockerfile
WORKDIR /storage

ENV USER=user
ENV PASS=password

EXPOSE 137/udp
EXPOSE 138/udp
EXPOSE 139/tcp
EXPOSE 445/tcp