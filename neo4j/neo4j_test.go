package neo4j

import "testing"

func TestConn(t *testing.T) {

	username := "neo4j"
	pwd := "minhaj"
	host := "localhost"
	port := "7687"
	store := NewNeo4jStore(username, pwd, host, port)
	_, err := store.Connect()
	if err != nil {
		t.Error(err)
	}

}
