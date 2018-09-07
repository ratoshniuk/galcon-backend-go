package main

import (
    "flag"
    "ws"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
    ws.StartServer(*addr)
    ws.StartClient(*addr)
}
