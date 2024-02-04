package model

import (
	"github.com/google/uuid"
	"log"
	"strings"
)

type Person struct {
	ID   string
	Name string
}

type Relationship struct {
	ID       string
	Name     string
	Relation string
}

type Connect struct {
	ID           string
	FirstPerson  string
	Connection   string
	SecondPerson string
}

var (
	validRelations = map[string]bool{"SON": true, "DAUGHTER": true, "FATHER": true, "MOTHER": true, "WIFE": true, "HUSBAND": true}
	pluralMap      = map[string]string{"sons": "SON", "daughters": "DAUGHTER", "wives": "WIFE"}
)

func ValidatePerson(person *Person) error {
	person.ID = uuid.New().String()

	return nil
}

func ValidateRelation(relation *Relationship) error {
	relation.Relation = strings.ToUpper(relation.Relation)

	if _, ok := validRelations[relation.Relation]; !ok {
		log.Fatalf("Unsupported relation type %v", relation.Relation)
	}

	relation.ID = uuid.New().String()

	return nil
}

func ValidateConnection(connect *Connect) error {
	connect.Connection = strings.ToUpper(connect.Connection)

	if _, ok := validRelations[connect.Connection]; !ok {
		log.Fatalf("Unsupported relation type %v", connect.Connection)
	}

	if connect.FirstPerson == connect.SecondPerson {
		log.Fatalf("Cannot create a relation on the same person")
	}

	return nil
}

func ValidateRelationType(connection string) (string, error) {
	relation := pluralMap[connection]

	if _, ok := validRelations[relation]; !ok {
		log.Fatalf("Unsupported relation type %v", connection)
	}

	return relation, nil
}
