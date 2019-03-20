package dua

import (
	"testing"

	"github.com/minhajuddinkhan/muntaha/neo4j"
)

func TestGetByEmotion(t *testing.T) {

	username := "neo4j"
	pwd := "minhaj"
	host := "localhost"
	port := "7687"
	store := neo4j.NewNeo4jStore(username, pwd, host, port)
	conn, err := store.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	s := NewService(conn)
	duas, err := s.GetByEmotion("weak")
	if err != nil {
		t.Error(err)
		return
	}

	if len(duas) == 0 {
		t.Fail()
	}
}
