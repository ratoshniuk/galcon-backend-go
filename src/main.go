package main

import (
    //"app"
    "bufio"
    //"flag"
    "fmt"
    "galcon-backend-go/messages/incoming"
    "galcon-backend-go/models"
    //"galcon-backend-go/rest"
    //"galcon-backend-go/ws"
    "net"
    "os"
    "strings"
)

func main() {

    //flag.Parse()
    //
    //context := &app.GlobalContext{}
    //context.Initialize()
    //context.SetRestAPI(&rest.Routes)
    //context.SetSocketAPI(&ws.Routes)
    //
    //context.Run()

    service := "127.0.0.1:8081"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    fmt.Println("it works!")
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {

        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }

        container:=populateTestDatabase()//for testing purposes only
        go handleRequest(container,conn)
    }

}
func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}

func handleRequest(container *models.GamesContainer,conn net.Conn) {
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

    parts := strings.Split(message,";")
    messageType, jsonBody := parts[0],parts[1]
    if messageType =="player_ready"{
        incoming.HandlePlayerReadyRequest(conn,container, jsonBody)
    }

}

func populateTestDatabase() *models.GamesContainer{
    container :=models.GamesContainer{}
    session:=models.GameSession{
        Id:1,
        Active:0,
        MaxPlayersCount:2,
    }
    session1:=models.GameSession{
        Id:2,
        Active:0,
        MaxPlayersCount:2,
    }
    session2:=models.GameSession{
        Id:3,
        Active:0,
        MaxPlayersCount:2,
    }
    models.AddGameSession(&container,&session)
    models.AddGameSession(&container,&session1)
    models.AddGameSession(&container,&session2)
    return &container
}