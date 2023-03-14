package cmd

import (
	"github.com/ha-dev/banking/app"
	"github.com/urfave/cli/v2"
)

// ────────────────────────────────────────────────────────────────────────────────

func (c *Commander) RunServer(ctx *cli.Context) error {
	appServer, err := app.NewApp()
	if err != nil {
		return err
	}

	err = appServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
