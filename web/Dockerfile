FROM php:apache

RUN apt update && apt upgrade

WORKDIR /usr/local/apache2/htdocs/
COPY ./public-html/ /var/www/html
