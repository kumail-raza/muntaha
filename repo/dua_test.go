package repo

import (
	"testing"

	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateRelationInRefAndDua(t *testing.T) {

	expected := `MATCH (m1:Dua{name: {Dua.name},translation: {Dua.translation},title: {Dua.title}}),(m2:Reference{name: {Reference.name}}) CREATE (m1)-[:REFERENCED_IN]->(m2)`
	ref := models.Reference{Name: "1235"}
	d := models.Dua{Arabic: "1", Title: "2", Translation: "3"}

	outputQ, args, err := CreateRelationInRefAndDua(ref.Name, d)

	assert.Nil(t, err, nil)
	assert.Equal(t, expected, outputQ)
	assert.Equal(t, ref.Name, args["Reference.name"])

}

func TestGetDuaByTitle(t *testing.T) {
	title := "123"
	expectedQuery := "MATCH (d:Dua{title: {Dua.title}}) RETURN d"

	outputQuery, args, err := GetDuaByTitle(title)
	assert.Nil(t, err)
	assert.Equal(t, expectedQuery, outputQuery)
	assert.Equal(t, title, args["Dua.title"])

}

func TestCreateRelationInEmoAndDua(t *testing.T) {

	e := models.Emotion{Name: "happy"}
	d := models.Dua{Arabic: "123"}
	expectedQ := `MATCH (e:Emotion{name: {emoName}}),(d:Dua{name: {duaName}}) CREATE (e)-[:RELATED]->(d)`
	q, args, err := CreateRelationInEmoAndDua(e, d)
	assert.Nil(t, err)
	assert.Equal(t, expectedQ, q)
	assert.Equal(t, e.Name, args["emoName"])

}

func TestCreateRelationInOriginAndDua(t *testing.T) {
	expected := `MATCH (m1:Dua{name: {Dua.name}}),(m2:Origin{type: {Origin.type},references: {Origin.references}}) CREATE (m1)-[:REFERENCED_IN]->(m2)`
	o := models.Origin{
		Type: "SomeType",
		References: []models.Reference{
			models.Reference{Name: "SomeReference"},
		},
	}
	d := models.Dua{
		Arabic: "dua",
	}

	q, args, err := CreateRelationInOriginAndDua(o, d)
	assert.Nil(t, err)
	assert.Equal(t, expected, q)
	assert.Equal(t, args["Dua.name"], d.Arabic)
	assert.Equal(t, args["Origin.type"], o.Type)
}

func TestGetAllDua(t *testing.T) {

	expected := `(d:Dua{})`
	q, args, err := GetAllDua()
	assert.Nil(t, err)
	assert.Equal(t, expected, q)
	assert.Nil(t, args)
}
