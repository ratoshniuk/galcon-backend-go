package outgoing

import (
	"encoding/json"
	"fmt"
	"net"
)

type PlayerKickedResponse struct {
	Reason string
}

func SendPlayerKickedResponse(response PlayerReadyResponse,connection net.Conn){

	jsonBody,_ :=json.Marshal(response)
	message:="player_kicked;"+string(jsonBody)+"\n"
	fmt.Fprintf(connection,message)


}
