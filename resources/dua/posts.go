package dua

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha/models"
	"github.com/minhajuddinkhan/muntaha/services/dua"
)

type CreateDuaRequest struct {
	Title       string             `json:"title,omitempty"`
	Arabic      string             `json:"arabic,omitempty"`
	Origin      string             `json:"origin,omitempty"`
	Translation string             `json:"translation,omitempty"`
	Emotions    []models.Emotion   `json:"emotions,omitempty"`
	References  []models.Reference `json:"references,omitempty"`
}

func (d *resource) CreateDua(req *restful.Request, res *restful.Response) {

	conn, err := d.Store.Connect()
	if err != nil {
		res.WriteError(http.StatusInternalServerError, fmt.Errorf("unable to connect to db  err: %v", err))
		return
	}
	defer conn.Close()

	var body CreateDuaRequest
	dec := json.NewDecoder(req.Request.Body)
	if err := dec.Decode(&body); err != nil {
		res.WriteError(http.StatusBadRequest, fmt.Errorf("unable to decode request body err: %v", err))
		return
	}
	reqDua := models.Dua{
		Arabic:      body.Arabic,
		Title:       body.Title,
		Translation: body.Translation,
	}
	reqOr := models.Origin{
		Type:       body.Origin,
		References: body.References,
	}

	ds := dua.NewService(conn)

	if err := ds.CreateDua(reqDua, body.Emotions, reqOr); err != nil {
		res.WriteError(http.StatusInternalServerError, fmt.Errorf("unable to create dua. err: %v", err))
		return
	}
	res.WriteHeaderAndEntity(http.StatusOK, "Dua created")
}
