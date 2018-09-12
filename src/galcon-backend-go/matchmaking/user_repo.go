package matchmaking

import (
    "fmt"
    "galcon-backend-go/cassandra"
    "github.com/gocql/gocql"
    "log"
)

type UserRepository interface {
    DDL(keyspace string) *string
    RegisterNew(u *User) (*User, error)
    Delete(id gocql.UUID) error
    RetrieveByID(id gocql.UUID) (*User, error)
    GetAll() (*[]User, error)
}

func UserRepoCassandraImpl(ctx *cassandra.CassandraContext) UserRepository {
    return &userRepo{
        persistence: ctx,
    }
}

func (repo *userRepo) DDL(keyspace string) *string {
    query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users (
    id uuid,
    rank int,
    PRIMARY KEY (id)
)`, keyspace)
    return &query
}

func (repo *userRepo) RegisterNew(u *User) (*User, error) {

    err := repo.persistence.KeyspaceSession.Query("insert into users (id, rank) values (?, ?)", u.ID, u.Rank).Exec()
    if err != nil {
        return nil, fmt.Errorf("error while register new user.cause: %s", err.Error())
    }
    return &User{ID: u.ID}, nil
}

func (repo *userRepo) Delete(id gocql.UUID) error {
    err := repo.persistence.KeyspaceSession.Query("delete from users where id = ?", id).Exec()
    if err != nil {
        return fmt.Errorf("error while deleting user.cause: %s", err.Error())
    }
    return nil
}

func (repo *userRepo) RetrieveByID(id gocql.UUID) (*User, error) {
    user := User{
        ID: id,
    }

    var rank int64

    if err := repo.persistence.KeyspaceSession.Query(`SELECT rank FROM users WHERE id = ? LIMIT 1`, id).Scan(&rank); err != nil {
        log.Fatal(err)
    }

    return &user, nil
}

func (repo *userRepo) GetAll() (*[]User, error) {
    var users []User

    var id gocql.UUID
    var rank int64

    iter := repo.persistence.KeyspaceSession.Query(`select id, rank FROM users`).Iter()
    for iter.Scan(&id, &rank) {
        users = append(users, User{
            ID: id,
            Rank: rank,
        })
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
    }

    return &users, nil
}
