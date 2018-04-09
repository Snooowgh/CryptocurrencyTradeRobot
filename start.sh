#!/bin/sh
if [ ! $1 ]
then
    rm nohup.out
    nohup go run main.go&
else
    go run main.go
fi
