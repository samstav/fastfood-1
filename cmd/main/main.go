package main

import (
	"log"
	"os"

	"github.com/jarosser06/fastfood/cmd"
	"github.com/mitchellh/cli"
)

const ffVersion = "0.1.2"

func main() {
	c := cli.NewCLI("fastfood", ffVersion)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"gen": func() (cli.Command, error) {
			return &cmd.Generator{}, nil
		},
		"build": func() (cli.Command, error) {
			return &cmd.Builder{}, nil
		},
		"new": func() (cli.Command, error) {
			return &cmd.Creator{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
