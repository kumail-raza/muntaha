package logger

import (
	restful "github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

func ApiLogger(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {

	chain.ProcessFilter(req, resp)
	logrus.Info(req.Request.Method, "\t", resp.StatusCode(), "\t", req.Request.RequestURI)
}
