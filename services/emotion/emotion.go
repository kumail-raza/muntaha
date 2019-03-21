package emotion

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/minhajuddinkhan/muntaha/models"
)

// Service Service
type Service interface {
	GetAll() ([]models.Emotion, error)
}

// NewService NewService
func NewService(conn bolt.Conn) Service {
	return &duaservice{Conn: conn}
}

type duaservice struct {
	Conn bolt.Conn
}

func (d *duaservice) GetAll() ([]models.Emotion, error) {

	nodes, _, _, err := d.Conn.QueryNeoAll(`MATCH (e:Emotion) RETURN e`, nil)
	if err != nil {
		return nil, err
	}

	emotions := make([]models.Emotion, len(nodes))
	for k, rawNode := range nodes {

		node := rawNode[0].(graph.Node)

		var emotion models.Emotion
		emotion.Name = node.Properties["name"].(string)
		emotions[k] = emotion
	}
	return emotions, nil

}
