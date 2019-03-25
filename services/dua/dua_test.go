package dua

import (
	"log"
	"testing"

	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/neo4j"
)

func getService() Service {

	username := "neo4j"
	pwd := "minhaj"
	host := "localhost"
	port := "7687"
	store := neo4j.NewNeo4jStore(username, pwd, host, port)
	conn, err := store.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return NewService(conn)
}

func TestGetByEmotion(t *testing.T) {

	duas, err := getService().GetByEmotion("weak")
	if err != nil {
		t.Error(err)
		return
	}

	if len(duas) == 0 {
		t.Fail()
	}
}

func TestCreateDuaWithQuran(t *testing.T) {

	d := models.Dua{Arabic: "a-test", Title: "title"}
	emos := []models.Emotion{models.Emotion{Name: "weak"}}
	o := models.Origin{Type: "Quran"}
	err := getService().CreateDua(d, emos, o)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestCreateDuaWithHadeeth(t *testing.T) {

	d := models.Dua{Arabic: "a-test", Title: "title"}
	emos := []models.Emotion{models.Emotion{Name: "weak"}}
	o := models.Origin{Type: "Hadeeth", References: []models.Reference{
		models.Reference{Name: "Ibn-Majah", RefNumber: "test"},
	}}

	err := getService().CreateDua(d, emos, o)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
