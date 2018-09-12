package matchmaking

import "fmt"

// TODO: please, use dummy for writing test scopes
type dummyRepo struct {
    persistence map[GameRoomID]*GameRoom
}

func GameRoomDummyImpl() GameRoomRepository {
    return &dummyRepo{
        persistence: make( map[GameRoomID]*GameRoom, 0),
    }
}

func (repo * dummyRepo) DDL(keyspace string) *string {
    return nil
}

func (repo * dummyRepo) RegisterNew(u *GameRoom) error {
    repo.persistence[u.ID] = u
    return nil
}

func (repo * dummyRepo) Delete(id GameRoomID) error {
    delete(repo.persistence, id)
    return nil
}

func (repo * dummyRepo) RetrieveById(id GameRoomID) (*GameRoom, error) {
    item:= repo.persistence[id]
    if item == nil {
        return nil, fmt.Errorf("game room with id %+v not found", id)
    }
    return item, nil
}
