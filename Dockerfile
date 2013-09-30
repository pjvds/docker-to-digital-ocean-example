FROM base
MAINTAINER Pieter Joost van de Sande <pj@born2code.net>

EXPOSE 80

RUN mkdir /var/www
ADD main /var/www/main

ENTRYPOINT ["/var/www/main"]
CMD ["--host=127.0.0.1", "--port=80"]