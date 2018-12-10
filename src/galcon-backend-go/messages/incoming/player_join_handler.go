package incoming

import (
	"encoding/json"
	"galcon-backend-go/container"
	"galcon-backend-go/models"
)

type PlayerJoinRequest struct {
	PlayerName string		`json:"player_name"`
}

func HandlePlayerJoinRequest(player *models.Player, container* container.GamesContainer, payload *json.RawMessage) {
	var request PlayerJoinRequest

	if json.Unmarshal(*payload, &request) != nil {
		panic("Unable to parse json!")
	}

	player.Login = request.PlayerName
	container.JoinQueue <- player
}

func HandlePlayerLeaveRequest(player *models.Player, container* container.GamesContainer, payload *json.RawMessage) {
	if player.Login == "" {
		return
	}

	container.LeaveQueue <- player
}