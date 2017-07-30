package cli

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/mitchellh/cli"
)

func Help(commands map[string]cli.CommandFactory) string {
	maxKeyLen := 0
	for key := range commands {
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
	}
	var buf bytes.Buffer
	buf.WriteString("usage: golodns [--version] <command> [args]\n\n")
	buf.WriteString("Available commands:\n")
	buf.WriteString(listCommands(commands, maxKeyLen))
	return buf.String()
}

func listCommands(commands map[string]cli.CommandFactory, maxKeyLen int) string {
	var buf bytes.Buffer
	keys := make([]string, 0, len(commands))
	for key := range commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		commandFunc, ok := commands[key]
		if !ok {
			// This should never happen since we JUST built the list of
			// keys.
			panic("command not found: " + key)
		}

		command, err := commandFunc()
		if err != nil {
			panic(fmt.Sprintf("command '%s' failed to load: %s", key, err))
		}

		key = fmt.Sprintf("%s%s", key, strings.Repeat(" ", maxKeyLen-len(key)))
		buf.WriteString(fmt.Sprintf("    %s    %s\n", key, command.Synopsis()))
	}

	return buf.String()
}
