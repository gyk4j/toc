#!/bin/sh

# Running as root. No need for `sudo`
# Reference: https://docs.docker.com/engine/install/centos/
yum update
yum install -y yum-utils
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

# Install networking tools
yum install -y net-tools samba-client

# Install specific version of Docker engine as CentOS Stream 9 will go 
# out-of-date, and newest Docker engine may no longer work in future.
# Specify the version that is used and tested to work.
yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
systemctl start docker
docker version

# Post-installation steps

# Manage Docker as a non-root user
groupadd docker
usermod -aG docker vagrant
newgrp docker

# Configure Docker to start on boot with systemd
systemctl enable docker.service
systemctl enable containerd.service

# Disable systemd-resolved on boot and use dnsmasq container instead
systemctl stop systemd-resolved
systemctl disable systemd-resolved
