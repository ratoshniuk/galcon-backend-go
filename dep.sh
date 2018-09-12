#!/bin/sh

echo "installing all deps.."

get_dep () {
echo "installing $1"
go get $2 $1
}

get_dep github.com/gorilla/websocket
get_dep github.com/gorilla/mux
get_dep github.com/jinzhu/gorm
get_dep github.com/hokaccha/go-prettyjson
get_dep github.com/gocql/gocql
get_dep github.com/satori/go.uuid

echo "...done"
