package neo4j

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Neo4j neo4j struct
type Neo4j interface {
	Connect() (bolt.Conn, error)
}

type neo4j struct {
	Username string
	Password string
	Host     string
	Port     string
}

// NewNeo4jStore instantiates a new store
func NewNeo4jStore(username, password, host, port string) Neo4j {
	return &neo4j{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

// Connect returns connection
func (n *neo4j) Connect() (bolt.Conn, error) {
	driver := bolt.NewDriver()
	conn := fmt.Sprintf("bolt://%s:%s@%s:%s", n.Username, n.Password, n.Host, n.Port)
	return driver.OpenNeo(conn)
}
