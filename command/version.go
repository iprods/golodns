package command

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type VersionCommand struct {
	HumanReadableVersion string
	UI                   cli.Ui
}

func (c *VersionCommand) Run(_ []string) int {
	c.UI.Output(fmt.Sprintf("golodns %s", c.HumanReadableVersion))
	return 0
}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Synopsis() string {
	return "Prints the current version"
}
