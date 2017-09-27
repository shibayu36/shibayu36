package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestBlogCommandRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &BlogCommand{outStream, errStream}

	args := []string{}

	if status := cli.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := "http://blog.shibayu36.org/"
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
