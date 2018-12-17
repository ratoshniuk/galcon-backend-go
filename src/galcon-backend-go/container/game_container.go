package container

import (
	"galcon-backend-go/messages/outgoing"
	"galcon-backend-go/models"
)

type GamesContainer struct {
	GameSessions map[int] *models.GameSession
	JoinQueue chan *models.Player
	LeaveQueue chan *models.Player
}

func NewGamesContainer() *GamesContainer {
	return &GamesContainer {
		JoinQueue: make(chan *models.Player),
		LeaveQueue: make(chan *models.Player),
		GameSessions: make(map[int] *models.GameSession),
	}
}

func (container *GamesContainer) Run() {
	for {
		select {
		case player := <-container.JoinQueue:
			session := container.findAvailableToJoinSession()
			session.AddPlayerToSession(player)
			freePlanet := session.GetFreePlanet()
			freePlanet.Player = player
			outgoing.NotifyPlayerJoined(session, player, freePlanet)
		case player := <-container.LeaveQueue:
			session := container.GetGameSessionById(player.SessionId)
			session.RemovePlayerFromSession(player)
			outgoing.NotifyPlayerLeft(session, player)
		}
	}
}

func (container *GamesContainer) findAvailableToJoinSession() *models.GameSession {
	for _, session := range container.GameSessions {
		if !session.IsFull() {
			return session
		}
	}

	newSession:= &models.GameSession{
		Id: len(container.GameSessions),
		MaxPlayersCount: 2,
		Players: make(map[int] *models.Player),
		Planets: container.generatePlanets(),
	}
	container.GameSessions[newSession.Id] = newSession
	return newSession
}

func (container *GamesContainer) generatePlanets() []*models.Planet {
	// TODO implement map generation
	return []*models.Planet{
		&models.Planet{Id: 1, Size: 6, Coordx: 1, Coordy: 1, Population: 45, Player:nil},
		&models.Planet{Id: 2, Size: 6, Coordx: 9, Coordy: 9, Population: 45, Player:nil},
		&models.Planet{Id: 3, Size: 4, Coordx: 2, Coordy: 8, Population: 30, Player:nil},
		&models.Planet{Id: 4, Size: 4, Coordx: 8, Coordy: 2, Population: 40, Player:nil},
		&models.Planet{Id: 5, Size: 4, Coordx: 4, Coordy: 4, Population: 50, Player:nil},
	}
}