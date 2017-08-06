package command

import (
	"flag"
	"fmt"
	"strings"
	"syscall"

	"github.com/iprods/golodns/resolver"
	"github.com/mitchellh/cli"
)

type InstallCommand struct {
	UI cli.Ui
}

func (c *InstallCommand) Run(args []string) int {
	if syscall.Getuid() != 0 {
		c.UI.Error("This command has to be run with sudo privileges as it writes to /etc/resolver")
		return 1
	}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }
	var d = f.String("domain", "localhost", "The name to resolve")
	var a = f.String("addr", "127.0.0.1", "The nameserver IP address")
	var p = f.String("port", "5300", "The nameserver port number")
	f.Parse(args)
	name := *d
	addr := *a
	port := *p
	r := resolver.Resolve{
		Path: "/etc/resolver",
	}
	domain, err := r.Install(name, addr, port)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(fmt.Sprintf("Successfully installed resolver for %s being handled by %s:%s", domain.Name, domain.IpAddress, domain.Port))
	c.UI.Output(fmt.Sprintf("golodns serve -addr %s -port %s -resolve_to [ip_to_resolve_to]", domain.IpAddress, domain.Port))
	return 0
}

func (c *InstallCommand) Help() string {
	helpText := `
Usage: sudo golodns install [flags]

  The install command installs the resolver for a specific domain. The following flags are possible:
   domain:		Which domain to map (default: localhost)
   address:		The nameserver IP address (default: 127.0.0.1)
   port:		The nameserver port number (default: 53)
	`
	return strings.TrimSpace(helpText)
}

func (c *InstallCommand) Synopsis() string {
	return "Install a resolver"
}
