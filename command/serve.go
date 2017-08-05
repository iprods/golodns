package command

import (
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/iprods/golodns/dns"
	"github.com/mitchellh/cli"
)

type ServeCommand struct {
	UI cli.Ui
}

func (c *ServeCommand) Run(args []string) int {
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }
	var addr = f.String("addr", "127.0.0.1", "The listening IP address")
	var port = f.String("port", "5300", "The listening port number")
	var resolveTo = f.String("resolve_to", "127.0.0.1", "The IP address to resolve to")
	f.Parse(args)
	resolveIp := net.ParseIP(*resolveTo)
	if resolveIp == nil {
		c.UI.Error("The resolving IP is not valid")
		return 1
	}
	if resolveIp.To4() == nil {
		c.UI.Error("The resolving IP is not an IPv4 address")
		return 1
	}
	server := dns.Server{
		Address:   *addr,
		Port:      *port,
		ResolveIp: resolveIp,
	}
	c.UI.Output(fmt.Sprintf("Listening on %s:%s, resolving to %s", *addr, *port, *resolveTo))
	c.UI.Error(server.Start().Error())

	return 0
}

func (c *ServeCommand) Help() string {
	helpText := `
Usage: golodns run [flags]

  The run command starts the resolver and listens to the defined address and port.
  The listening address and port can be provided by passing the "-addr" and "-port" flags
	e.g. -addr=127.0.0.1 -port 5300 (default).
  The resolving address can be provided by passing the "-resolv_to" flag
	e.g. -resolve_to=127.0.0.1 (default).
	`
	return strings.TrimSpace(helpText)
}

func (c *ServeCommand) Synopsis() string {
	return "Start the local DNS resolver"
}
