package matchmaking

import (
    "galcon-backend-go/cassandra"
    "github.com/gocql/gocql"
    "testing"
)

func TestBaseCrudDummy(t *testing.T) {

    var userRepo = UserRepoDummyImpl()

    testUser := &User{ID: gocql.TimeUUID()}
    u, err := userRepo.RegisterNew(testUser)
    if err != nil {
        t.Errorf("Error while creating new user %+v", err)
    }

    u2, err := userRepo.RetrieveByID(u.ID)
    if err != nil {
        t.Errorf("Error while retrieving user %+v", err)
    }
    if u2 != u {
        t.Errorf("Error while fetching same user from repo")
    }
}

func TestBaseCrudCassadra(t *testing.T) {

    var cass = cassandra.CassandraContext{}
    cass.Init("127.0.0.1:9042", "test")

    defer cass.Stop(false)

    var userRepo = UserRepoCassandraImpl(&cass)

    cass.PerformDDL(userRepo.DDL)

    testUser := User{ID: gocql.TimeUUID(), Rank: int64(1)}

    u, err := userRepo.RegisterNew(&testUser)
    if err != nil {
        t.Errorf("Error while creating new user %+v", err)
    }

    defer userRepo.Delete(u.ID)

    u2, err := userRepo.RetrieveByID(u.ID)
    if err != nil {
       t.Errorf("Error while retrieving user %+v", err)
    }

    if u2.Rank != u.Rank {
       t.Errorf("Error while fetching same user from repo")
    }
}
