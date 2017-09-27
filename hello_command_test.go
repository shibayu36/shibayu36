package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelloCommandRun(t *testing.T) {
	{
		// error occurs without args
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &HelloCommand{outStream, errStream}

		if status := cli.Run([]string{}); status != ExitCodeError {
			t.Errorf("hello command without args should occur error")
		}
		if !strings.Contains(errStream.String(), "Please specify name") {
			t.Errorf("expected %q to eq %q", errStream.String(), "Please specify name")
		}
	}

	{
		// hello command success
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &HelloCommand{outStream, errStream}

		if status := cli.Run([]string{"shibayu36"}); status != ExitCodeOK {
			t.Errorf("hello command doesn't exit correctly")
		}
		if !strings.Contains(outStream.String(), "Hello shibayu36") {
			t.Errorf("hello command output is invalid")
		}
	}
}
