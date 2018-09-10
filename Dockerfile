FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./src /go/src
ADD fresh.conf /go/src
ADD dep.sh /go/src
ADD revision.txt /go/src
ADD version.txt /go/src

WORKDIR /go/src

RUN \
    sh ./dep.sh

RUN go build -o galcon .

CMD ./galcon

EXPOSE 8080
