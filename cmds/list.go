package cmds

import (
	"fmt"

	"github.com/mitchellh/go-spdx"
	"github.com/urfave/cli/v2"
)

var List = &cli.Command{
	Name:  "list",
	Usage: "List licenses",
	Action: func(c *cli.Context) error {
		list, err := spdx.List()
		if err != nil {
			return err
		}
		for _, v := range list.Licenses {
			fmt.Println(v.ID)
		}
		return nil
	},
}
