package main

import "fmt"

// BlogCommand is CLI struct for `blog` sub-command
type BlogCommand struct{}

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
	fmt.Printf("http://blog.shibayu36.org/\n")
	return 0
}
