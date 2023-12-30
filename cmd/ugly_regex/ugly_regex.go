package main

import (
	"fmt"
	"os"

	regex "github.com/12yanogden/regex/internal"
	"golang.design/x/clipboard"
)

func main() {
	cmd := "ugly_regex"
	args := os.Args[1:]
	err := clipboard.Init()
	var pretty string

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(args) == 1 {
		fmt.Println("No need to pass arguments. '" + cmd + "' scans your clipboard")
		pretty = args[0]
	} else if len(args) > 1 {
		fmt.Printf("%s: expected 0 or 1 arguments, but %d arguments given\n", cmd, len(args))
		os.Exit(1)
	} else {
		pretty = string(clipboard.Read(clipboard.FmtText))
	}

	clipboard.Write(clipboard.FmtText, []byte(regex.Ugly(pretty)))
}
