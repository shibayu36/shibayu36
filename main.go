package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("shibayu36", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"hello": func() (cli.Command, error) {
			return &HelloCommand{}, nil
		},
		"blog": func() (cli.Command, error) {
			return &BlogCommand{}, nil
		},
	}
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
