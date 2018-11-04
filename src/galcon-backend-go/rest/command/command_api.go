package info

import (
	"app"
	"encoding/json"
	"galcon-backend-go/rest/common"
	"log"
	"net/http"
)

func CommandHandler(ctx *app.GlobalContext, rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var t CommandRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t)

	switch t.Command {

	case "player_ready":
		log.Println("Got 'player_ready' command")
		args := t.Arguments.(map[string]interface{})
		playerName := args["player_name"].(string)
		common.RespondJSON(rw, http.StatusOK, &PlayerReadyResponse{
			Success:   true,
			Arguments: map[string]string{"player_name": playerName},
		})

	case "kicked":
		log.Println("Got 'kicked' command")
		args := t.Arguments.(map[string]interface{})
		playerName := args["player_name"].(string)
		reason := args["reason"].(string)
		common.RespondJSON(rw, http.StatusOK, &PlayerKickedResponse{
			Success:   true,
			Arguments: map[string]string{"player_name": playerName, "reason": reason},
		})

	}

}
