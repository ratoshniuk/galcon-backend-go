package container

import "galcon-backend-go/models"

func (container *GamesContainer) GetGameSessionById(id int) *models.GameSession  {
	session := container.GameSessions[id]
	if session == nil {
		panic("Session with id " + string(id) + " not found!")
	}
	return session
}

func (container *GamesContainer) getPlayerById(sessionId int, playerId int) *models.Player{
	gameSession := container.GetGameSessionById(sessionId)
	player := gameSession.Players[playerId]
	if player == nil {
		panic("player not found!")
	}
	return player
}

func (container *GamesContainer) GetPlayersFromSession(id int) map[int] *models.Player {
	return container.GetGameSessionById(id).Players
}

func (container *GamesContainer) SetPlayerReady(sessionId int, playerId int ) *models.Player{
	player := container.getPlayerById(sessionId, playerId)
	player.Ready=true
	return player
}

func (container *GamesContainer) UpdateSessionStatus(sessionId int){
	gameSession := container.GetGameSessionById(sessionId)
	if gameSession.MaxPlayersCount == len(gameSession.Players) {
		for _, player := range gameSession.Players{
			if !player.Ready {
				return
			}
		}
		gameSession.Active=true
	}
}
