package main

import (
	"os"

	"github.com/autotelic/contentful-fetcher/cli/cmd"
)

func main() {
	ctl := cmd.NewDefaultRootCommand(os.Args[0])
	ctl.Execute()
}
