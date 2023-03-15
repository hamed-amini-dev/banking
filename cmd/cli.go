package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

type Commander struct {
	app cli.App
}

/*
New Commander initialize for getting args from console
configure command args console
assign command args to objects
*/
func New() (*Commander, error) {
	var commander Commander

	//init commands
	commander.app = cli.App{
		Commands: []*cli.Command{
			{
				Name:    "startServer",
				Aliases: []string{"s", "serve"},
				Usage:   "start http server",
				Action:  commander.RunServer,
			},
		},
	}

	return &commander, nil
}

/*
Run Commander object previously initialized
*/
func (c *Commander) Run() error {
	return c.app.Run(os.Args)
}
