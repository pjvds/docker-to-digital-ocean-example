FROM base
MAINTAINER Pieter Joost van de Sande <pj@born2code.net>

EXPOSE 5000

RUN mkdir /var/www
ADD main /var/www/main

ENTRYPOINT ["/var/www/main"]
CMD ["--host=0.0.0.0", "--port=5000"]