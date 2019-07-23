FROM docker:latest

RUN apk add --no-cache --update make git docker bash curl

WORKDIR /

COPY ./maas.sh /usr/local/bin/maas.sh

RUN chmod +x /usr/local/bin/maas.sh

ENTRYPOINT [ "maas.sh" ]