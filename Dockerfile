FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./src /go/src
WORKDIR /go/src

RUN \
	go get github.com/gorilla/websocket && \
	go get github.com/gorilla/mux && \
	go get github.com/jinzhu/gorm && \
	go get github.com/hokaccha/go-prettyjson && \
	go get github.com/denisenkom/go-mssqldb

RUN go build

ADD fresh.conf /go/src

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh -c fresh.conf; \
	fi


EXPOSE 8080