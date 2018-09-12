package matchmaking

type GameRoomRepository interface {
    DDL(keyspace string) *string
    RegisterNew(u *GameRoom) error
    Delete(id GameRoomID) error
    RetrieveById(id GameRoomID) (*GameRoom, error)

    //RetrieveParticipants(id GameRoomID) (*[]UserId, error)
    //RegisterParticipant(id GameRoomID, userId UserId) error
    //RemoveParticipant(id GameRoomID, userId UserId) error
}


//func (repo *dummyRepo) RetrieveParticipants(id GameRoomID) (*[]UserId, error) {
//    room, err := repo.RetrieveById(id)
//    if err != nil {
//        return nil, err
//    }
//    return &room.Participants, nil
//}
//
//func (repo *dummyRepo) RegisterParticipant(id GameRoomID, userId UserId) error {
//    room, err := repo.RetrieveById(id)
//    if err != nil {
//        return err
//    }
//    old := room.Participants
//    old = append(old, userId)
//    room.Participants = old
//    repo.persistence[id] = room
//    return nil
//}
//
//func (repo *dummyRepo) RemoveParticipant(id GameRoomID, userId UserId) error {
//    room, err := repo.RetrieveById(id)
//    if err != nil {
//        return err
//    }
//
//    new := make([]UserId, 0)
//    for _, u := range room.Participants {
//        if u != userId {
//            new = append(new, u)
//        }
//    }
//    room.Participants = new
//    repo.persistence[id] = room
//    return nil
//}
