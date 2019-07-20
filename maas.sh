#! /usr/bin/env bash

git_url=$1

rm -rf gitmaas | sed -e 's/^/[maaslog]: /'

git clone $git_url gitmaas | sed -e 's/^/[maaslog]: /'

cd gitmaas

make | sed -e 's/^/[maaslog]: /'