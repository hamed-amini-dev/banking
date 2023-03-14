package main

import (
	"github.com/ha-dev/banking/cmd"
	"github.com/ha-dev/banking/pkg/config"
)

func main() {

	err := config.InitConfig(".")
	if err != nil {
		panic(err)
	}

	commander, err := cmd.New()
	if err != nil {
		panic(err)
	}
	if err := commander.Run(); err != nil {
		panic(err)
	}

}
