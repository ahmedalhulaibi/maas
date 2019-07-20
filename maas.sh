#! /usr/bin/env bash

git_url=$1

rm -rf gitmaas

git clone $git_url gitmaas

cd gitmaas

make