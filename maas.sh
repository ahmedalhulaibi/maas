#! /usr/bin/env bash

git_url=$1
echo "git URL:" $git_url

makeArgs="${@:2}"
echo "Make args: ${makeArgs[@]}"

rm -rf gitmaas | sed -e 's/^/[maaslog]: /'

git clone $git_url gitmaas | sed -e 's/^/[maaslog]: /'

cd gitmaas

make $makeArgs | sed -e 's/^/[maaslog]: /'