package command

import (
	"fmt"
	"strings"

	"github.com/iprods/golodns/resolver"
	"github.com/mitchellh/cli"
)

type ListCommand struct {
	UI cli.Ui
}

func (c *ListCommand) Run(args []string) int {
	r := resolver.Resolve{
		Path: "/etc/resolver",
	}
	domains, err := r.List()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	for _, domain := range domains {
		entryMessage := fmt.Sprintf("%s:%s handles domain %s", domain.IpAddress, domain.Port, domain.Name)
		if domain.Managed {
			entryMessage += " and is managed by golodns"
		}
		c.UI.Info(entryMessage)
	}
	return 0
}

func (c *ListCommand) Help() string {
	helpText := `
Usage: golodns list

  The list command return all registered resolvers or an error of none are present.
	`
	return strings.TrimSpace(helpText)
}

func (c *ListCommand) Synopsis() string {
	return "List the available resolvers"
}
