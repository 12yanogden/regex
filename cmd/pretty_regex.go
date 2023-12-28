package main

import (
	"fmt"
	"os"

	regex "github.com/12yanogden/regex/internal"
	"golang.design/x/clipboard"
)

func main() {
	cmd := "pretty_regex"
	args := os.Args[1:]
	err := clipboard.Init()

	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		fmt.Println(cmd + ": no arguments given")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("%s: expected 1 argument, but %d arguments given\n", cmd, len(args))
		os.Exit(1)
	}

	clipboard.Write(clipboard.FmtText, []byte(regex.Pretty(args[0])))
}
