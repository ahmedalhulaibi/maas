#! /usr/bin/env bash

git_url=$1
allArgs="${@:2}"

rm -rf gitmaas | sed -e 's/^/[maaslog]: /'

git clone $git_url gitmaas | sed -e 's/^/[maaslog]: /'

cd gitmaas

make $allArgs | sed -e 's/^/[maaslog]: /'