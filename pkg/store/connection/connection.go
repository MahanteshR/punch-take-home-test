package connection

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/punch-test/pkg/model"
	"log"
)

type Connection struct {
	db *sql.DB
}

func New(db *sql.DB) *Connection {
	return &Connection{db: db}
}

const (
	selectConnection = "SELECT FROM" +
		" connection WHERE EXISTS(SELECT 1 FROM connection WHERE ((first_person=$1 AND second_person=$2) OR (first_person=$2 AND second_person=$1)))"

	countConnections = "SELECT COUNT(relationship) FROM" +
		" connection WHERE relationship=$1 AND second_person=$2"

	getConnection = "SELECT first_person FROM" +
		" connection WHERE relationship=$1 AND second_person=$2"

	insertConnection = "INSERT INTO" + " connection(id,first_person,relationship,second_person) VALUES($1,$2,$3,$4)"
)

func (c Connection) AddConnection(connection *model.Connect) error {
	var exists bool

	_ = c.db.QueryRow(selectConnection, connection.FirstPerson, connection.SecondPerson).Scan(&exists)

	if exists {
		log.Fatalf("A connection between %v and %v already exists", connection.FirstPerson, connection.SecondPerson)
	}

	_, err := c.db.Exec(insertConnection, uuid.New().String(), connection.FirstPerson, connection.Connection, connection.SecondPerson)
	if err != nil {
		return err
	}

	log.Printf("Relationship has been added, [%v] is the [%v] of [%v]", connection.FirstPerson, connection.Connection, connection.SecondPerson)
	return nil
}

func (c Connection) CountConnections(relation, name string) (int, error) {
	var count int

	err := c.db.QueryRow(countConnections, relation, name).Scan(&count)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	return count, nil
}

func (c Connection) GetConnection(relation, name string) error {
	var firstPerson string

	err := c.db.QueryRow(getConnection, relation, name).Scan(&firstPerson)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	log.Printf("%v is %v of %v", firstPerson, relation, name)
	return nil
}
