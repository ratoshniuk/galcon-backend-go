package models

import (
	"net"
	"time"
)

type Message struct {
	Type string
	Payload interface{}
}

type Planet struct {
	Id int
	Size int
	Population int
	Coordx int
	Coordy int
}

type Group struct {
	Id int
	Amount int
	Coordx int//only if we implement redirect
	Coordy int//only if we implement redirect
	TargetPlanet Planet
	SourcePlanet Planet
	SourceGroup *Group//only if we implement redirect
	ArrivalTime	time.Time
}

type Player struct {
	Id int
	SessionId int
	Connection *net.Conn
	Login string
	Ready bool
}

type GameSession struct {
	Id int
	Active bool
	MaxPlayersCount int
	Planets []*Planet
	Groups []*Group
	Players map[int] *Player
}

func (session *GameSession) IsFull() bool {
	return len(session.Players) == session.MaxPlayersCount
}

func (session *GameSession) AddPlayerToSession(player *Player)  bool {
	if session.IsFull() {
		return false
	}

	player.Id = len(session.Players)
	player.SessionId = session.Id
	session.Players[player.Id] = player
	return true
}

func (session *GameSession) RemovePlayerFromSession(player *Player) {
	if session.Active {
		player.Ready = false
	} else {
		delete(session.Players, player.Id)
	}

	(*player.Connection).Close()
}
