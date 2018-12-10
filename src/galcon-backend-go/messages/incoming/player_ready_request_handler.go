package incoming

import (
	"encoding/json"
	"galcon-backend-go/container"
	"galcon-backend-go/messages/outgoing"
	"galcon-backend-go/models"
)

type PlayerReadyRequest struct {
	SessionId int
	PlayerId  int
}

func HandlePlayerReadyRequest(player *models.Player, container* container.GamesContainer, payload *json.RawMessage) {
	var requestBody PlayerReadyRequest
	err := json.Unmarshal(*payload, &requestBody)

	if err != nil {
		panic("Wrong json!")
	}

	updatedPlayer := container.SetPlayerReady(requestBody.SessionId,requestBody.PlayerId)
	container.UpdateSessionStatus(requestBody.SessionId)

	players := container.GetPlayersFromSession(requestBody.SessionId)
	for _, player := range players {
		if player.Ready {
			response := &outgoing.PlayerReadyResponse{Login: updatedPlayer.Login}
			msg := &models.Message{
				Type: outgoing.PlayerReadyMessageType,
				Payload: response,
			}
			outgoing.SendJsonResponse(msg, player.Connection)
		}
	}
}
