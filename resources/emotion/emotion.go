package emotion

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha/logger"
	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/neo4j"
	"github.com/minhajuddinkhan/muntaha/services/emotion"
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
		Path("/emotion").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("/").
			Filter(logger.ApiLogger).
			To(r.GetAll).
			Doc("gets all emotions").
			Writes([]models.Emotion{}).
			Returns(http.StatusOK, "OK", []models.Emotion{}),
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
	es := emotion.NewService(conn)
	emotions, err := es.GetAll()
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	response.WriteEntity(emotions)
}
