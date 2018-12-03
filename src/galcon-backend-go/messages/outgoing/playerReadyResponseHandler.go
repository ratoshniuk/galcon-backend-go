package outgoing

import (
	"encoding/json"
	"fmt"
	"net"
)

type PlayerReadyResponse struct {
	Login string
}

func SendPlayerReadyResponse(response PlayerReadyResponse,connection net.Conn){

	jsonBody,_ :=json.Marshal(response)
	message:="player_ready;"+string(jsonBody)+"\n"
	fmt.Fprintf(connection,message)


}