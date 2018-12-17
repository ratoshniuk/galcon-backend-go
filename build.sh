#!/usr/bin/env bash

echo "installing deps..."
sh ./dep.sh
echo "running tests..."

go test -timeout 600s -v ./src/galcon-backend-go/...

buildCmd=""
echo $buildCmd

if [[ "$1" == "windows" ]]
then
    buildCmd="GOOS=windows go build -o"
elif [[ "$1" == "linux" ]]
then
    buildCmd"CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o"
elif [[ "$1" == "macos" ]]
then
    buildCmd="GOOS=darwin go build -o"
else
	buildCmd="go build -o"
fi

cmd="$buildCmd galcon ./src"
eval ${cmd}
