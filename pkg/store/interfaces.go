package store

import (
	"github.com/punch-test/pkg/model"
)

type Person interface {
	AddPerson(person *model.Person) error
}

type Relation interface {
	AddRelation(relation *model.Relationship) error
}

type Connection interface {
	AddConnection(connect *model.Connect) error
	CountConnections(relation, name string) (int, error)
	GetConnection(relation, name string) error
}
