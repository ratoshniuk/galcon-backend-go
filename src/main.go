package main

import (
    "app"
    "flag"
    "galcon-backend-go/rest"
    "galcon-backend-go/ws"
)

func main() {

    flag.Parse()

    context := &app.GlobalContext{}
    context.Initialize()
    context.SetRestAPI(&rest.Routes)
    context.SetSocketAPI(&ws.Routes)

    context.Run()

}
