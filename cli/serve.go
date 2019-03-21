package cli

import (
	"fmt"
	"net/http"

	"github.com/minhajuddinkhan/muntaha/resources"

	"github.com/emicklei/go-restful"
	"github.com/minhajuddinkhan/muntaha"
	"github.com/minhajuddinkhan/muntaha/neo4j"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var port string

// Serve serves an http server on a port
func Serve(conf muntaha.Configuration) cli.Command {
	return cli.Command{
		Name: "serve",
		Action: func(c *cli.Context) error {
			store := neo4j.NewNeo4jStore(conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port)

			wsContainer := restful.NewContainer()
			mgr := resources.NewResourceManager(store)
			mgr.SpawnAPIContainer(wsContainer)

			logrus.Infof("listening on port  %s", conf.HttpPort)
			return http.ListenAndServe(fmt.Sprintf(":%s", conf.HttpPort), wsContainer)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "port, p",
				Usage:       "port to spawn an http server over",
				Destination: &port,
			},
		},
	}

}
