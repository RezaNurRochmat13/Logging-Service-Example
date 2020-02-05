FROM frolvlad/alpine-glibc

RUN apk add --no-cache bash

ADD main /

EXPOSE 8081

CMD ["/main"]