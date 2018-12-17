package main

import (
	"bufio"
	"encoding/json"
	"galcon-backend-go/container"
	"galcon-backend-go/messages/incoming"
	"galcon-backend-go/models"
	"log"
	"net"
	"os"
)

const (
	//ConnHost = "localhost"
	ConnPort = "8081"
	ConnType = "tcp"
)

type handler func(*models.Player, *container.GamesContainer, *json.RawMessage)

var RequestHandlers = map[string]handler{
	"player_ready": incoming.HandlePlayerReadyRequest,
	"join":         incoming.HandlePlayerJoinRequest,
	"leave":        incoming.HandlePlayerLeaveRequest,
	"send_ships":   incoming.HandleSendShipsRequest,
}

func main() {
	listener, err := net.Listen(ConnType, ":"+ConnPort)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	log.Println("Listening on " + ":" + ConnPort)

	gameContainer := container.NewGamesContainer()
	go gameContainer.Run()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		player := models.Player{
			Connection: &conn,
		}
		go handleRequest(gameContainer, &player)
	}
}

func handleRequest(container *container.GamesContainer, player *models.Player) {

	for {
		msgType, payload, err := readMessage(player.Connection)
		if err {
			log.Printf("Connection closed for player with id: %d and session id: %d", player.Id, player.SessionId)
			return
		}

		requestHandler := RequestHandlers[msgType]
		if requestHandler != nil {
			requestHandler(player, container, payload)
		}
	}
}

func readMessage(conn *net.Conn) (string, *json.RawMessage, bool) {
	request, err := bufio.NewReader(*conn).ReadString('\n')
	if err != nil {
		log.Println("Error reading:", err.Error())
		return "", nil, true
	}

	log.Printf("Received message: %s", request)

	var payload json.RawMessage
	msg := models.Message{
		Payload: &payload,
	}
	if err := json.Unmarshal([]byte(request), &msg); err != nil {
		log.Fatal(err)
	}
	return msg.Type, &payload, false
}
