FROM alpine:latest

RUN apk add --update make git jq zip

WORKDIR /

COPY ./maas.sh /usr/local/bin/maas.sh

RUN chmod +x /usr/local/bin/maas.sh

ENTRYPOINT [ "maas.sh" ]