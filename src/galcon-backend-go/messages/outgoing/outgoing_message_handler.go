package outgoing

import (
	"encoding/json"
	"galcon-backend-go/models"
	"net"
)

const (
	JoinAcceptedMessageType = "join_accepted"
	PlayerJoinedMessageType = "player_joined"
	PlayerReadyMessageType = "player_ready"
	PlayerLeftMessageType = "player_left"
	PlayerKickedMessageType = "player_kicked"
)

type PlayerReadyResponse struct {
	Login string				`json:"login"`
}

type PlayerKickedResponse struct {
	Reason string				`json:"reason"`
}

type JoinAcceptedResponse struct {
	PlayerId int				`json:"player_id"`
	SessionId int				`json:"session_id"`
	StartingPlanetId int		`json:"starting_planet_id"`
	Planets []*models.Planet	`json:"planets"`
	GrowthRate float64			`json:"population_growth_rate"`
}

type PlayerJoinedResponse struct {
	PlayerName string			`json:"name"`
	StartingPlanetId int		`json:"starting_planet_id"`
}

type PlayerLeftResponse struct {
	PlayerName string			`json:"player_name"`
}

func SendJsonResponse(message *models.Message, connection *net.Conn) {
	jsonBody, _ := json.Marshal(message)
	(*connection).Write(jsonBody)
}

func NotifyPlayerLeft(session *models.GameSession, leftPlayer *models.Player)  {
	msg := &models.Message {
		Type: PlayerLeftMessageType,
		Payload: &PlayerLeftResponse {
			PlayerName: leftPlayer.Login,
		},
	}
	notifyAllExceptSender(msg, session, leftPlayer)
}

func NotifyPlayerJoined(session *models.GameSession, joinedPlayer *models.Player) {
	notifyJoinedPlayer(session, joinedPlayer)
	notifyOtherPlayers(session, joinedPlayer)
}

func notifyJoinedPlayer(session *models.GameSession, joinedPlayer *models.Player) {
	joinAcceptedMsg := models.Message{
		Type: JoinAcceptedMessageType,
		Payload: &JoinAcceptedResponse {
			PlayerId: joinedPlayer.Id,
			SessionId: session.Id,
			Planets: session.Planets,

			// change after planet generation logic will be implemented
			StartingPlanetId: joinedPlayer.Id + 3,
			GrowthRate: 6.316,
		},
	}
	SendJsonResponse(&joinAcceptedMsg, joinedPlayer.Connection)
}

func notifyOtherPlayers(session *models.GameSession, joinedPlayer *models.Player) {
	playerJoinedMsg := &models.Message {
		Type: PlayerJoinedMessageType,
		Payload: &PlayerJoinedResponse {
			PlayerName: joinedPlayer.Login,
			StartingPlanetId: joinedPlayer.Id + 3,
		},
	}
	notifyAllExceptSender(playerJoinedMsg, session, joinedPlayer)
}

func notifyAllExceptSender(msg *models.Message, session *models.GameSession, sender *models.Player) {
	for _, player := range session.Players {
		if player.Id != sender.Id {
			SendJsonResponse(msg, player.Connection)
		}
	}
}