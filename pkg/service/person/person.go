package person

import (
	"github.com/punch-test/pkg/model"
	"github.com/punch-test/pkg/store"
)

type Person struct {
	store store.Person
}

func New(person store.Person) *Person {
	return &Person{store: person}
}

func (p *Person) AddPerson(person *model.Person) error {
	err := model.ValidatePerson(person)
	if err != nil {
		return err
	}

	err = p.store.AddPerson(person)
	if err != nil {
		return err
	}

	return nil
}
