package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zakuro9715/license/cmds"
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
		Commands: []*cli.Command{
			cmds.List,
		},
	}

	app.RunAndExitOnError()
}
