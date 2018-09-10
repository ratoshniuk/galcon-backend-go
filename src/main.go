package main

import (
    "app"
    "config"
    "flag"
    "galcon-backend-go/rest"
    "galcon-backend-go/ws"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

    flag.Parse()

    context := &app.Context{}
    context.Initialize(config.GetConfig())
    context.SetRestAPI(&rest.Routes)
    context.SetSocketAPI(&ws.Routes)

    context.Run(*addr)

}
