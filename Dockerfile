FROM ubuntu:latest

RUN apt-get update && apt-get install make git jq zip -y

WORKDIR /

COPY ./maas.sh /usr/local/bin/maas.sh

RUN chmod +x /usr/local/bin/maas.sh

ENTRYPOINT [ "maas.sh" ]