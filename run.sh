#!/bin/sh

docker run --rm -ti -p 8080:8080 $(docker build -t galcon-daemon -q ./)
