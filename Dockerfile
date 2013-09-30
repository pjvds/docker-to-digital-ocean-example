FROM ubuntu
MAINTAINER Pieter Joost van de Sande <pj@born2code.net>

EXPOSE 80

RUN mkdir /var/www
ADD main /var/www/main

ENTRYPOINT /var/wwww/main
