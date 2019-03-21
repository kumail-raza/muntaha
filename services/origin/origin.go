package origin

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/minhajuddinkhan/muntaha/models"
)

// Service Service
type Service interface {
	GetAll() ([]models.Origin, error)
}

// NewService NewService
func NewService(conn bolt.Conn) Service {
	return &service{Conn: conn}
}

type service struct {
	Conn bolt.Conn
}

func (s *service) GetAll() ([]models.Origin, error) {
	nodes, _, _, err := s.Conn.QueryNeoAll(`MATCH (o:Origin) RETURN o`, nil)
	if err != nil {
		return nil, err
	}
	origins := make([]models.Origin, len(nodes))
	for k, rawNode := range nodes {
		node := rawNode[0].(graph.Node)
		origins[k].Type = node.Properties["type"].(string)
	}
	return origins, nil
}
