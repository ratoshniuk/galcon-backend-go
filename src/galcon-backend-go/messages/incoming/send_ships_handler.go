package incoming

import (
	"encoding/json"
	"galcon-backend-go/container"
	"galcon-backend-go/messages/outgoing"
	"galcon-backend-go/models"
	"log"
	"math"
	"time"
)

type SendShipsRequest struct {
	FromPlanetId int `json:"from"`
	ToPlanetId int `json:"to"`
}

func HandleSendShipsRequest(player *models.Player, container* container.GamesContainer, payload *json.RawMessage) {
	var requestBody SendShipsRequest
	err := json.Unmarshal(*payload, &requestBody)

	if err != nil {
		panic("Wrong json!")
	}

	gameSession := container.GetGameSessionById(player.SessionId)
	if gameSession.Active == false {
		log.Println("Session is not active")
		return
	}

	sourcePlanet:= gameSession.GetPlanetById(requestBody.FromPlanetId)

	if sourcePlanet == nil {
		log.Printf("Source planet not found")
		return
	}

	if sourcePlanet.Player == nil || sourcePlanet.Player.Id != player.Id {
		log.Printf("User is not owner of source planet")
		return
	}

	targetPlanet := gameSession.GetPlanetById(requestBody.ToPlanetId)

	if targetPlanet == nil {
		log.Printf("Target planet not found")
		return
	}

	xDistance := math.Pow(float64(sourcePlanet.Coordx - targetPlanet.Coordx), 2)
	yDistance := math.Pow(float64(sourcePlanet.Coordy - targetPlanet.Coordy), 2)
	distanceBetweenPlanets := math.Sqrt(xDistance + yDistance)

	group := &models.Group{
		Id: len(gameSession.Groups),
		Amount: 2,
		SourcePlanet:*sourcePlanet,
		TargetPlanet:*targetPlanet,
		ArrivalTime: time.Now().Add(time.Second * time.Duration(distanceBetweenPlanets)),
		Player: *player,
	}
	gameSession.Groups = append(gameSession.Groups, group)

	sourcePlanet.Population = sourcePlanet.Population / 2
	// TODO call populationSync when implemented (?).

	for _, player := range gameSession.Players {
		response := &outgoing.ShipsSentResponse{
			GroupId:group.Id,
			FromPlanetId:group.SourcePlanet.Id,
			ToPlanetId:group.TargetPlanet.Id,
			Amount:group.Amount,
			ArrivalTimestamp:group.ArrivalTime.Unix(),
		}

		msg := &models.Message{
			Type: outgoing.ShipsSentResponseMessageType,
			Payload: response,
		}

		outgoing.SendJsonResponse(msg, player.Connection)
	}


	time.AfterFunc(time.Duration(group.ArrivalTime.Sub(time.Now()).Seconds()) * time.Second, func() {
		log.Printf("Ships arrived");
		//TODO call shipsArrived handler when implemented
	})
}