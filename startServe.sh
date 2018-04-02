#!/bin/bash

stop(){
    if [ -f "./runtime/rpc-srv.pid" ]; then
        echo "stopping last run"
        cat ./runtime/rpc-srv.pid | xargs kill
        rm -f ./runtime/rpc-srv.pid
    fi

    if [ -f "./runtime/rpc-srv.out" ];then
        rm -f ./runtime/rpc-srv.out
    fi
}

stop

rm -f rpc-srv

echo "build server"

go build -o rpc-srv serve/serve.go

echo "starting server"

if [ ! -f "./rpc-srv" ]; then
    echo "rpc-srv is not valid"
else
    nohup ./rpc-srv > ./runtime/rpc-srv.out 2>&1 &
    echo $! > ./runtime/rpc-srv.pid
    sleep 3
    if [ -s ./runtime/rpc-srv.out ]; then
        cat ./runtime/rpc-srv.out
    else
        echo "rpc server is running"
    fi
fi

