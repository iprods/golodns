package cli

import (
	"os"

	"github.com/iprods/golodns/command"
	"github.com/iprods/golodns/version"
	"github.com/mitchellh/cli"
)

// Command factory providing all available commands
func Commands() map[string]cli.CommandFactory {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}
	return map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				UI: ui,
			}, nil
		},
		"install": func() (cli.Command, error) {
			return &command.InstallCommand{
				UI: ui,
			}, nil
		},
		"serve": func() (cli.Command, error) {
			return &command.ServeCommand{
				UI: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				HumanReadableVersion: version.HumanReadableVersion(),
				UI:                   ui,
			}, nil
		},
	}
}
