package main

import (
	"github.com/urfave/cli/v2"
)

var (
	name    = "license"
	version = "latest"
)

func main() {
	app := &cli.App{
		Name:    name,
		Usage:   "license tool",
		Version: version,
	}

	app.RunAndExitOnError()
}
