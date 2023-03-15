package main

import (
	"github.com/ha-dev/banking/cmd"
	"github.com/ha-dev/banking/pkg/config"
)

/*
Load Configuration File
Create New Commander package for getting args console
Run Commander package with args inputted
*/
func main() {

	//load configuration with viper and init it
	err := config.InitConfig(".")
	if err != nil {
		panic(err)
	}

	//commander initialize for getting args from console
	commander, err := cmd.New()
	if err != nil {
		panic(err)
	}

	//run commander
	if err := commander.Run(); err != nil {
		panic(err)
	}

}
