package dua

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/querybuilder"
)

// Service Service
type Service interface {
	GetAll() ([]models.Dua, error)
	GetByEmotion(name string) ([]models.Dua, error)
	CreateDua(dua models.Dua, emo []models.Emotion, org models.Origin) error
}

// NewService NewService
func NewService(conn bolt.Conn) Service {
	return &duaservice{Conn: conn}
}

type duaservice struct {
	Conn bolt.Conn
}

const (
	getDuaByTitle = "MATCH(d:Dua{ title: {title} }) RETURN d"
)

func (d *duaservice) GetAll() ([]models.Dua, error) {
	query, args := querybuilder.GetAllDua()
	nodes, _, _, err := d.Conn.QueryNeoAll(query, args)
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
		`MATCH(e:Emotion{name: {name}}),(d:Dua)
		MATCH (e)-[:RELATED]->(d)
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

func (d *duaservice) CreateDua(dua models.Dua, emos []models.Emotion, org models.Origin) error {

	query, args := querybuilder.GetDuaByTitle(dua.Title)
	nodes, _, _, err := d.Conn.QueryNeoAll(query, args)
	if err != nil {
		return err
	}
	if len(nodes) > 0 {
		return fmt.Errorf("dua already exists with this title")
	}

	query, args = querybuilder.CreateDua(dua, org)
	if _, err = d.Conn.ExecNeo(query, args); err != nil {
		return err
	}

	switch org.Type {
	case "Hadeeth":
		for _, ref := range org.References {
			query, args := querybuilder.CreateRelationInRefAndDua(ref.Name, models.Dua{Title: dua.Title})
			if _, err := d.Conn.ExecNeo(query, args); err != nil {
				return err
			}
		}
		break
	case "Quran":
		q, args := querybuilder.CreateRelationInOriginAndDua(org, models.Dua{Title: dua.Title})
		if _, err := d.Conn.ExecNeo(q, args); err != nil {
			return err
		}
		break
	}

	for _, e := range emos {
		q, args := querybuilder.CreateRelationInEmoAndDua(e, models.Dua{Title: dua.Title})
		if _, err := d.Conn.ExecNeo(q, args); err != nil {
			return err
		}
	}

	return nil
}
