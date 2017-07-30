package main // import "github.com/iprods/golodns"

import (
	"os"

	"github.com/iprods/golodns/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
