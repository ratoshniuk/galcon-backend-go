package matchmaking

import (
    "github.com/gocql/gocql"
)

type User struct {
    Rank int64
    ID gocql.UUID// TODO : with uuid replace
}

type Status int

const (
    StatusPending = iota
    StatusReady
    StatusPLaying
    StatusFinished
)

type GameRoomID gocql.UUID // TOOD uuid replace

type GameRoom struct {
    ID GameRoomID
    Status Status
    //TTL int64
    Dimension [][]int
}

type GameRoomParticipants struct {
    Participants []gocql.UUID
}

func NewAwaiting(creator *gocql.UUID) *GameRoom {
    return &GameRoom{
        ID : GameRoomID(gocql.TimeUUID()),
    }
}


