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
			outgoing.NotifyPlayerJoined(session, player)
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
	}
	container.GameSessions[newSession.Id] = newSession
	return newSession
}