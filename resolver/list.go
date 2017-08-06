package resolver

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Generate a list of installed resolvers by checking and evaluating the files in a given directory.
func (r *Resolve) List() ([]Domain, error) {
	var domains = []Domain{}
	files, err := ioutil.ReadDir(r.Path)
	if err != nil {
		return nil, errors.New("No resolvers found.")
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		domain, err := r.loadDomainResolverInfo(name)
		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

func (r *Resolve) loadDomainResolverInfo(name string) (Domain, error) {
	path := r.Path + "/" + name
	f, _ := os.Open(path)
	defer f.Close()
	fscan := bufio.NewScanner(f)
	var ipAddress string
	port := "53"
	managed := false
	for fscan.Scan() {
		line := strings.TrimSpace(fscan.Text())
		if line == "# Managed by golodns" {
			managed = true
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, "#")
		line = strings.TrimSpace(parts[0])
		ws := regexp.MustCompile(`\s+`)
		parts = ws.Split(line, -1)
		if len(parts) != 2 {
			return Domain{}, errors.New(fmt.Sprintf("Unexpected line. Found %d parts.", len(parts)))
		}
		entryType := parts[0]
		entryValue := parts[1]
		if entryType == "nameserver" {
			ipAddress = entryValue
			continue
		}
		if entryType == "port" {
			port = entryValue
			continue
		}
	}
	domain := Domain{
		Name:      name,
		IpAddress: ipAddress,
		Port:      port,
		Managed:   managed,
	}
	return domain, nil
}
