package resolver

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Uninstall the file in the resolver directory
func (r *Resolve) Uninstall(name string) (Domain, error) {
	domain := Domain{
		Name: name,
	}
	path := r.Path
	resFile := path + "/" + name
	// Check if the file is managed
	f, err := os.Open(resFile)
	defer f.Close()
	if err != nil {
		return Domain{}, errors.New(fmt.Sprintf("Domain resolver for %s does not exist.", name))
	}
	fScan := bufio.NewScanner(f)
	managed := false
	for fScan.Scan() {
		line := strings.TrimSpace(fScan.Text())
		if line == "# Managed by golodns" {
			managed = true
			break
		}
	}
	if !managed {
		return Domain{}, errors.New(fmt.Sprintf("Domain %s is not managed.", name))
	}
	err = os.Remove(resFile)
	if err != nil {
		return Domain{}, err
	}
	return domain, nil
}
