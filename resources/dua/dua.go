package dua

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha/logger"
	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/neo4j"
)

// Resource dua http resource
type Resource interface {
	WebService() *restful.WebService
	GetAll(request *restful.Request, response *restful.Response)
}

// NewResource creates a new dua http resource
func NewResource(store neo4j.Neo4j) Resource {
	return &resource{Store: store}
}

type resource struct {
	Store neo4j.Neo4j
}

func (d *resource) WebService() *restful.WebService {

	ws := new(restful.WebService)
	ws.
		Path("/dua").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("/").
			Filter(logger.ApiLogger).
			To(d.GetAll).
			Doc("gets all duas").
			Writes([]models.Dua{}).
			Returns(http.StatusOK, "OK", []models.Dua{}),
	)

	ws.Route(
		ws.GET("/{emotion}").
			Filter(logger.ApiLogger).
			To(d.GetByEmotion).
			Doc("gets all duas by emotion").
			Writes([]models.Dua{}).
			Returns(http.StatusOK, "OK", []models.Dua{}),
	)

	return ws
}
