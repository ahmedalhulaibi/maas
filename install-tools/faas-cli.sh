#!/bin/sh
PLATFORM=$1

if [[ "$PLATFORM" = "Alpine" ]]; then
    curl -sSL https://cli.openfaas.com | sh
elif [[ "$PLATFORM" = "Debian" ]]; then
    curl -sSL https://cli.openfaas.com | sudo -E sh
elif [[ "$PLATFORM" = "Darwin" ]]; then
    if [ -x "$(command -v brew)" ]; then
        brew install faas-cli
    else
        curl -sSL https://cli.openfaas.com | sudo -E sh
    fi
fi