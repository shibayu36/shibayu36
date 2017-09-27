package main

import (
	"fmt"
	"io"
)

// BlogCommand is CLI struct for `blog` sub-command
type BlogCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message
func (c *BlogCommand) Synopsis() string {
	return "Print shibayu36 blog information"
}

// Help is content of help for this command
func (c *BlogCommand) Help() string {
	return "Usage: shibayu36 blog"
}

// Run is main method
func (c *BlogCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "http://blog.shibayu36.org/")
	return 0
}
