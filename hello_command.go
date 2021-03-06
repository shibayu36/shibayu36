package main

import (
	"fmt"
	"io"
)

// HelloCommand is CLI struct for `hello` sub-command
type HelloCommand struct{
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message
func (c *HelloCommand) Synopsis() string {
	return "Say Hello"
}

// Help is content of help for this command
func (c *HelloCommand) Help() string {
	return "Usage: shibayu36 hello (name)"
}

// Run is main method
func (c *HelloCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(c.ErrStream, "Please specify name")
		return 1
	}

	fmt.Fprintf(c.OutStream, "Hello %s\n", args[0])
	return 0
}
