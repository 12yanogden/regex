package main

import (
	"fmt"
	"os"

	"github.com/12yanogden/errors"
	"github.com/12yanogden/regex/internal/regex"
	"golang.design/x/clipboard"
)

func main() {
	args := os.Args[1:]
	err := clipboard.Init()

	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		errors.Scream("no arguments given")
	} else if len(args) > 1 {
		errors.Scream(fmt.Sprintf("expected 1 argument, but %d arguments given.", len(args)))
	}

	clipboard.Write(clipboard.FmtText, []byte(regex.Pretty(args[0])))
}
