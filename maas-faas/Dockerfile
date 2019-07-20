FROM openfaas/classic-watchdog:0.13.4 as watchdog

FROM docker:latest

RUN mkdir -p /home/app

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

# Custom to maas-faas
USER root
RUN apk add --no-cache --update make git docker bash
COPY ./maas-faas.sh /usr/bin/maas-faas.sh

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="maas-faas.sh"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
