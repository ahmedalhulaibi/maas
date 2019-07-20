#! /usr/bin/env bash

readarray lineArgs < /dev/stdin
declare -a allArgs
i=0
for arg in $lineArgs
do
    echo "$i: $arg"
    allArgs[$i]="$arg"
    let i=$i+1
done
echo "All args: ${allArgs[@]}"

git_url=${allArgs[0]}
echo "git URL:" $git_url

makeArgs=${allArgs[@]:1}
echo "Make args: ${makeArgs[@]}"

rm -rf gitmaas | sed -e 's/^/[maaslog]: /'

git clone $git_url gitmaas | sed -e 's/^/[maaslog]: /'

cd gitmaas

make $makeArgs | sed -e 's/^/[maaslog]: /'
