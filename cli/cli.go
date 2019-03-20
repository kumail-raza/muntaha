package cli

import (
	"os"

	"github.com/minhajuddinkhan/muntaha"

	"github.com/urfave/cli"
)

const (
	// Version version of cli application
	Version = "0.1.0"
)

// Run runs the cli application
func Run(conf muntaha.Configuration) error {
	app := cli.NewApp()
	app.Version = Version
	app.Commands = []cli.Command{
		Serve(conf),
	}
	return app.Run(os.Args)
}
