package relation

import (
	"database/sql"
	"github.com/punch-test/pkg/model"
	"log"
)

type Relation struct {
	db *sql.DB
}

func New(db *sql.DB) *Relation {
	return &Relation{db: db}
}

const (
	insertQuery = "INSERT INTO" + " relation(id,name,type) VALUES ($1,$2,$3)"
)

func (r Relation) AddRelation(relation *model.Relationship) error {
	_, err := r.db.Exec(insertQuery, relation.ID, relation.Name, relation.Relation)
	if err != nil {
		return err
	}

	log.Printf("relaton [%v] has been successfully added to the person [%v]", relation.Relation, relation.Name)
	return nil
}
