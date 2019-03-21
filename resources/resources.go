package resources

import (
	"github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha/neo4j"
	"github.com/minhajuddinkhan/muntaha/resources/dua"
	"github.com/minhajuddinkhan/muntaha/resources/emotion"
	"github.com/minhajuddinkhan/muntaha/resources/origin"
	"github.com/minhajuddinkhan/muntaha/resources/reference"
)

type ResourceManager interface {
	SpawnAPIContainer(c *restful.Container)
}

func NewResourceManager(store neo4j.Neo4j) ResourceManager {
	return &resManager{Store: store}
}

type resManager struct {
	Store neo4j.Neo4j
}

func (mgr *resManager) SpawnAPIContainer(c *restful.Container) {

	c.Add(dua.NewResource(mgr.Store).WebService())
	c.Add(emotion.NewResource(mgr.Store).WebService())
	c.Add(origin.NewResource(mgr.Store).WebService())
	c.Add(reference.NewResource(mgr.Store).WebService())

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{
			"Content-Type", "Accept",
			"Access-Control-Allow-Origin", "*",
		},
	}
	c.Filter(cors.Filter)

}
