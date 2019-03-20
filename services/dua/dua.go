package dua

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/minhajuddinkhan/muntaha/models"
)

// Service Service
type Service interface {
	GetAll() ([]models.Dua, error)
	GetByEmotion(name string) ([]models.Dua, error)
}

// NewService NewService
func NewService(conn bolt.Conn) Service {
	return &duaservice{Conn: conn}
}

type duaservice struct {
	Conn bolt.Conn
}

func (d *duaservice) GetAll() ([]models.Dua, error) {
	nodes, _, _, err := d.Conn.QueryNeoAll(
		`MATCH(d:Dua) RETURN d`, nil)
	if err != nil {
		return nil, err
	}
	duas := make([]models.Dua, len(nodes))
	for k, rawNode := range nodes {
		node := rawNode[0].(graph.Node)

		var dua models.Dua
		dua.Arabic = node.Properties["name"].(string)
		dua.Title = node.Properties["title"].(string)
		dua.Translation = node.Properties["translation"].(string)

		duas[k] = dua
	}

	return duas, nil
}
func (d *duaservice) GetByEmotion(name string) ([]models.Dua, error) {

	cypher := map[string]interface{}{"name": name}
	nodes, _, _, err := d.Conn.QueryNeoAll(
		`MATCH(d:Dua),(e:Emotion) WHERE e.name = {name}
		RETURN d`, cypher)

	if err != nil {
		return nil, err
	}

	duas := make([]models.Dua, len(nodes))
	for k, rawNode := range nodes {
		node := rawNode[0].(graph.Node)

		var dua models.Dua
		dua.Arabic = node.Properties["name"].(string)
		dua.Title = node.Properties["title"].(string)
		dua.Translation = node.Properties["translation"].(string)

		duas[k] = dua
	}

	return duas, nil
}
