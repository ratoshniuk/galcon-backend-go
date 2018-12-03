package models

func AddGameSession(container *GamesContainer,gameSession *GameSession) {
	container.GameSessions= append(container.GameSessions, gameSession)

}
func AddPlayer(container *GamesContainer,sessionId int, player *Player){
	session:=GetGameSessionById(container,sessionId)
	if len(session.Players)<session.MaxPlayersCount {
		player.Ready=0
		GetGameSessionById(container,sessionId).Players=append(session.Players, player)
	}

}

func GetGameSessionById(container *GamesContainer,id int) *GameSession  {
	for _,session :=range container.GameSessions{
		if session.Id==id {
			return session
		}
	}
	panic("Session with id "+string(id)+" not found!")
}

func getPlayerById(container *GamesContainer, sessionId int, playerId int) *Player{
	gameSession := GetGameSessionById(container,sessionId)
	for _,player := range gameSession.Players{
		if player.Id == playerId{
			return player
		}
	}
	panic("player not found!")
}

func GetListOfPlayersBySession(container *GamesContainer,id int) []*Player {
	for _, session := range container.GameSessions {
		if session.Id == id {
			return session.Players
		}
	}
	panic("Session with id " + string(id) + " not found!")
}

func SetPlayerReady(container *GamesContainer ,sessionId int, playerId int ) *Player{
	player := getPlayerById(container,sessionId,playerId)
	player.Ready=1
	return player
}

func UpdateSessionStatus(container *GamesContainer ,sessionId int){
	gameSession := GetGameSessionById(container,sessionId)
	if gameSession.MaxPlayersCount == len(gameSession.Players) {

		for _, player := range gameSession.Players{
			if player.Ready !=1{
				return
			}
		}
	}else {
		return
	}
	gameSession.Active=1


}
