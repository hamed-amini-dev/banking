package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

type Commander struct {
	app cli.App
}

// ────────────────────────────────────────────────────────────────────────────────
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

// ────────────────────────────────────────────────────────────────────────────────
func (c *Commander) Run() error {
	return c.app.Run(os.Args)
}

// ────────────────────────────────────────────────────────────────────────────────
