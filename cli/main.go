package cli

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func Run(args []string) int {
	return RunCustom(args, Commands())
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {
	for _, arg := range args {
		if arg == "--" {
			break
		}
		if arg == "-v" || arg == "--version" {
			args = []string{"version"}
			break
		}
	}
	commandsInclude := make([]string, len(commands))
	for k, _ := range commands {
		commandsInclude = append(commandsInclude, k)
	}
	cli := &cli.CLI{
		Name:         "golodns",
		Args:         args,
		Commands:     commands,
		HelpFunc:     cli.FilteredHelpFunc(commandsInclude, Help),
		Autocomplete: true,
	}
	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
