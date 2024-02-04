package connection

import (
	"github.com/punch-test/pkg/model"
	"github.com/punch-test/pkg/store"
	"strings"
)

type Connection struct {
	store store.Connection
}

func New(connection store.Connection) *Connection {
	return &Connection{store: connection}
}

func (c *Connection) AddConnection(connect *model.Connect) error {
	err := model.ValidateConnection(connect)
	if err != nil {
		return err
	}

	err = c.store.AddConnection(connect)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) CountConnections(relation, name string) (*int, error) {
	relation = strings.ToUpper(relation)

	relation, err := model.ValidateRelationType(relation)
	if err != nil {
		return nil, err
	}

	count, err := c.store.CountConnections(relation, name)
	if err != nil {
		return nil, err
	}

	countPtr := &count

	return countPtr, nil
}

func (c *Connection) GetConnection(relation, name string) error {
	err := c.store.GetConnection(strings.ToUpper(relation), name)
	if err != nil {
		return err
	}

	return nil
}
