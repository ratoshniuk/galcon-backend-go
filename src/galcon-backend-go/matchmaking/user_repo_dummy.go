package matchmaking

import (
	"fmt"
	"github.com/gocql/gocql"
)

type userRepoDummy struct {
	persistence map[gocql.UUID]*User
}

func UserRepoDummyImpl() UserRepository {
	return &userRepoDummy{
		persistence: make(map[gocql.UUID]*User, 0),
	}
}

func (repo *userRepoDummy) DDL(keyspace string) *string {
	return nil
}

func (repo *userRepoDummy) RegisterNew(u *User) (*User, error) {
	repo.persistence[u.ID] = u
	return u, nil
}

func (repo *userRepoDummy) Delete(id gocql.UUID) error {
	delete(repo.persistence, id)
	return nil
}

func (repo *userRepoDummy) RetrieveByID(id gocql.UUID) (*User, error) {
	user := repo.persistence[id]
	if user == nil {
		return nil, fmt.Errorf("user with id %+v does not exist", id)
	}
	return user, nil
}

func (repo *userRepoDummy) GetAll() (*[]User, error) {
	all := make([]User, 0)
	for _, user := range repo.persistence {
		all = append(all, *user)
	}
	return &all, nil
}
