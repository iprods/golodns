package command

import (
	"flag"
	"fmt"
	"strings"
	"syscall"

	"github.com/iprods/golodns/resolver"
	"github.com/mitchellh/cli"
)

type UninstallCommand struct {
	UI cli.Ui
}

func (c *UninstallCommand) Run(args []string) int {
	if syscall.Getuid() != 0 {
		c.UI.Error("This command has to be run with sudo privileges as it writes to /etc/resolver")
		return 1
	}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }
	var n = f.String("domain", "localhost", "The name to resolve")
	f.Parse(args)
	r := resolver.Resolve{
		Path: "/etc/resolver",
	}
	name := *n
	domain, err := r.Uninstall(name)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(fmt.Sprintf("Successfully uninstalled resolver for %s.\n", domain.Name))
	return 0
}

func (c *UninstallCommand) Help() string {
	helpText := `
Usage: sudo golodns uninstall [flags]

  The uninstall command removes the resolver for a specific domain. The following flags are possible:
   domain:		Which domain to uninstall (default: localhost)
	`
	return strings.TrimSpace(helpText)
}

func (c *UninstallCommand) Synopsis() string {
	return "Uninstall a resolver"
}
