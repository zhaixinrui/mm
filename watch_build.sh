#!/bin/bash

#go build mm.go

DIR=$(dirname "$0")

function build(){
    echo $1|egrep '\.go$'
    if [ 0 -eq $? ];then
        go build *.go >$DIR/buildresult 2>&1
        cat $DIR/buildresult
        rm -f $DIR/buildresult
        echo "build finish"
    # else
        # echo "$1 change, ignore"
    fi
}

if [ -z $1 ];then
    # echo "begin to watch directory: $DIR"
    fswatch $DIR -e .git | xargs -n1 $0
else
    build $1
fi

# while True; do
#     fswatch $DIR -e .git -i *.go -1
#     build
# done