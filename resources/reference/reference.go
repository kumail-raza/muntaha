package reference

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha/logger"
	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/neo4j"
	"github.com/minhajuddinkhan/muntaha/services/reference"
)

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

func (r *resource) WebService() *restful.WebService {
	ws := new(restful.WebService)

	ws.
		Path("/reference").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("/").
			Filter(logger.ApiLogger).
			To(r.GetAll).
			Doc("gets all reference").
			Writes([]models.Emotion{}).
			Returns(http.StatusOK, "OK", []models.Reference{}),
	)
	return ws
}

// GetAll gets all emotions
func (r *resource) GetAll(request *restful.Request, response *restful.Response) {

	conn, err := r.Store.Connect()
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	rs := reference.NewService(conn)
	references, err := rs.GetAll()
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	response.WriteEntity(references)
}
