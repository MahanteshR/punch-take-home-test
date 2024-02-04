package relation

import (
	"github.com/punch-test/pkg/model"
	"github.com/punch-test/pkg/store"
)

type Relation struct {
	store store.Relation
}

func New(relation store.Relation) *Relation {
	return &Relation{store: relation}
}

func (r *Relation) AddRelation(relation *model.Relationship) error {
	err := model.ValidateRelation(relation)
	if err != nil {
		return err
	}

	err = r.store.AddRelation(relation)
	if err != nil {
		return err
	}

	return nil
}
