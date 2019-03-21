package reference

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/minhajuddinkhan/muntaha/models"
)

// Service Service
type Service interface {
	GetAll() ([]models.Reference, error)
}

// NewService NewService
func NewService(conn bolt.Conn) Service {
	return &service{Conn: conn}
}

type service struct {
	Conn bolt.Conn
}

func (s *service) GetAll() ([]models.Reference, error) {
	nodes, _, _, err := s.Conn.QueryNeoAll(`MATCH(r:Reference) RETURN r`, nil)
	if err != nil {
		return nil, err
	}
	references := make([]models.Reference, len(nodes))
	for k, rawNode := range nodes {
		node := rawNode[0].(graph.Node)
		references[k].ID = node.Properties["id"].(int64)
		references[k].Name = node.Properties["name"].(string)
	}
	return references, nil
}
