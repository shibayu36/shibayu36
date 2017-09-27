package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/cli"
)

// CLI is struct for switch stdout / stderr
type CLI struct {
	outStream, errStream io.Writer
}

// Run method
func (c *CLI) Run(args []string) int {
	cl := cli.NewCLI("shibayu36", Version)
	cl.Args = args[1:]

	cl.Commands = map[string]cli.CommandFactory{
		"hello": func() (cli.Command, error) {
			return &HelloCommand{c.outStream, c.errStream}, nil
		},
		"blog": func() (cli.Command, error) {
			return &BlogCommand{c.outStream, c.errStream}, nil
		},
	}

	exitStatus, err := cl.Run()

	if err != nil {
		fmt.Fprintf(c.errStream, "Failed to execute: %s\n", err.Error())
	}

	return exitStatus
}

func main() {
	c := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(c.Run(os.Args))
}
