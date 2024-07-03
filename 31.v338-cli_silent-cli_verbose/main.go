package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		silent  bool
		verbose bool
	)

	flag.BoolVar(&silent, "s", false, "silent mode")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	var (
		appW   io.Writer = os.Stdout
		dbgW   io.Writer = os.Stdout
		withLf           = func(s string) string {
			return fmt.Sprintf("%s\n", s)
		}
	)

	if silent {
		dbgW = io.Discard
	}

	if verbose {
		io.WriteString(appW, withLf("zero"))
	}

	io.WriteString(appW, withLf("first"))
	io.WriteString(dbgW, withLf("second"))

	if verbose {
		io.WriteString(appW, withLf("fourth"))
	}
}
