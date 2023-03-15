package main

import (
	"os"
	"testing"

	"github.com/ha-dev/banking/cmd"
	"github.com/ha-dev/banking/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("Config test", func(t *testing.T) {
		err := config.InitConfig(".")
		assert.Nil(t, err)
	})

	t.Run("run main test with args", func(t *testing.T) {
		os.Args = []string{"s"}
		commander, err := cmd.New()
		assert.Nil(t, err)
		err = commander.Run()
		assert.Nil(t, err)
	})

}
