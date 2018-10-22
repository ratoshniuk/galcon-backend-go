package matchmaking

//
import (
	//"fmt"
	"github.com/gocql/gocql"
	//"log"
)

type UserRepository interface {
	DDL(keyspace string) *string
	RegisterNew(u *User) (*User, error)
	Delete(id gocql.UUID) error
	RetrieveByID(id gocql.UUID) (*User, error)
	GetAll() (*[]User, error)
}
