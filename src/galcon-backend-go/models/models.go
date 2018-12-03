package models

import (
	"net"
	"time"
)

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
	Connection net.Conn
	Login string
	Ready int
}

type GameSession struct {
	Id int
	Active int
	MaxPlayersCount int
	Planets []*Planet
	Groups []*Group
	Players []*Player
}

type GamesContainer struct {
	GameSessions []*GameSession
}


