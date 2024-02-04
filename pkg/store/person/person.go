package person

import (
	"database/sql"
	"github.com/punch-test/pkg/model"
	"log"
)

type Person struct {
	db *sql.DB
}

func New(db *sql.DB) *Person {
	return &Person{db: db}
}

const (
	insertQuery = "INSERT INTO" + " person(id, name) VALUES ($1, $2)"
)

func (p *Person) AddPerson(person *model.Person) error {
	_, err := p.db.Exec(insertQuery, person.ID, person.Name)
	if err != nil {
		return err
	}

	log.Printf("person [%v] has been successfully added to the family tree", person.Name)
	return nil
}
