package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func Test_VersionCommand_is_a_cli_command(t *testing.T) {
	var _ cli.Command = &VersionCommand{}
}
