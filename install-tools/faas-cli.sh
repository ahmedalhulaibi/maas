#!/bin/bash
PLATFORM=$1
echo "install-tools/faas-cli.sh running"
if [ -x "$(command -v faas-cli)" ]; then
    echo "faas-cli already installed"
    exit 0
fi

if [ "$PLATFORM" = "Alpine" ]; then
    curl -sSL https://cli.openfaas.com > faas-cli-install.sh
    chmod +x ./faas-cli-install.sh
    ./faas-cli-install.sh
elif [ "$PLATFORM" = "Debian" ]; then
    curl -sSL https://cli.openfaas.com > faas-cli-install.sh
    chmod +x ./faas-cli-install.sh
    ./faas-cli-install.sh
elif [ "$PLATFORM" = "Darwin" ]; then
    if [ -x "$(command -v brew)" ]; then
        brew install faas-cli
    else
        curl -sSL https://cli.openfaas.com > faas-cli-install.sh
        chmod +x ./faas-cli-install.sh
        ./faas-cli-install.sh
    fi
fi