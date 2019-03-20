package dua

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha/services/dua"
)

func (d *resource) GetAll(request *restful.Request, response *restful.Response) {
	conn, err := d.Store.Connect()
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	defer conn.Close()
	ds := dua.NewService(conn)
	duas, err := ds.GetAll()
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	response.WriteEntity(duas)

}
func (d *resource) GetByEmotion(request *restful.Request, response *restful.Response) {
	conn, err := d.Store.Connect()
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	defer conn.Close()
	emo := request.PathParameter("emotion")
	ds := dua.NewService(conn)
	duas, err := ds.GetByEmotion(emo)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	response.WriteEntity(duas)

}
