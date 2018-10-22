package matchmaking

import (
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
