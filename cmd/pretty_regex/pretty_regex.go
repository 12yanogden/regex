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
	var ugly string

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(args) == 1 {
		fmt.Println("No need to pass arguments. '" + cmd + "' scans your clipboard")
		ugly = args[0]
	} else if len(args) > 1 {
		fmt.Printf("%s: expected 0 or 1 arguments, but %d arguments given\n", cmd, len(args))
		os.Exit(1)
	} else {
		ugly = string(clipboard.Read(clipboard.FmtText))
	}

	clipboard.Write(clipboard.FmtText, []byte(regex.Pretty(ugly)))
}
