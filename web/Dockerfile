FROM php:apache

RUN apt update && apt upgrade 

# Just use a bind mount
# apt install -y cifs-utils
# RUN mkdir /media/storage
# COPY .smbcredentials /root/.smbcredentials
# RUN chmod 600 /root/.smbcredentials
# RUN echo //file/Data /media/storage cifs credentials=/root/.smbcredentials 0 0 >> /etc/fstab
# CMD ["mount", "/media/storage"]

COPY ./public-html/ /var/www/html

WORKDIR /var/www/html
