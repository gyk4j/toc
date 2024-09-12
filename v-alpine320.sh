#!/bin/sh

# Reference:
#   - https://wiki.alpinelinux.org/wiki/Docker
#   - https://docs.genesys.com/Documentation/System/latest/DDG/InstallationofDockeronAlpineLinux
apk update && apk upgrade

# Install login tools and OpenRC
apk add util-linux-login openrc 

# Install networking tools
apk add net-tools samba-client

# Install Docker
apk add docker docker-engine docker-cli containerd docker-cli-buildx docker-cli-compose

# Post-installation steps

# Manage Docker as a non-root user
addgroup -S docker
adduser root docker
adduser vagrant docker
newgrp docker

rc-service docker start
rc-service docker status
docker version

# Configure Docker to start on boot with OpenRC
rc-update add docker boot
# rc-update add containerd boot

