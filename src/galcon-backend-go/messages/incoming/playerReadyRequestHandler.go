package incoming

import (
	"awesomeProject/messages/outgoing"
	"awesomeProject/models"
	"encoding/json"
	"net"
)

type PlayerReadyRequest struct {
	SessionId int
	PlayerId  int
}

func HandlePlayerReadyRequest(connection  net.Conn,container* models.GamesContainer,jsonString string) {

	var requestBody PlayerReadyRequest
	err := json.Unmarshal([]byte(jsonString), &requestBody)
	if err == nil {

		//Remove next 2 rows when player_joined is implemented!
		newPlayer:=models.Player{Id:requestBody.PlayerId,Login:"Player1",Connection:connection}//for testing purposes only
		models.AddPlayer(container,requestBody.SessionId,&newPlayer)//for testing purposes only

		updatedPlayer:=models.SetPlayerReady(container,requestBody.SessionId,requestBody.PlayerId)
		models.UpdateSessionStatus(container,requestBody.SessionId)

		players := models.GetListOfPlayersBySession(container,requestBody.SessionId)
		for _,player := range players{
			if player.Ready != -1 {
				response := outgoing.PlayerReadyResponse{Login: updatedPlayer.Login}
				outgoing.SendPlayerReadyResponse(response, player.Connection)
			}
		}

	} else {
		panic("Wrong json!")
	}
}
