package main

import (
    "flag"
    "fmt"
    "github/com/galcon-backend-go/ws"
    "os/exec"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {

    res := exec.Command("echo $HEROKU_API_SECRET")

    fmt.Printf("%+v", res)

    ws.StartServer(*addr)
    ws.StartClient(*addr)
}
