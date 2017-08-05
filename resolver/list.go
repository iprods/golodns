package resolver

import (
	"io/ioutil"
	"os"
	"bufio"
	"strings"
	"regexp"
)

type Domain struct {
	Name      string
	IpAddress string
	Port      string
	Managed   bool
}

func (r *Resolve) List() ([]Domain		, error) {
	var domains = []Domain{}
	files, err := ioutil.ReadDir(r.Path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			// TODO Panic
		}
		name := file.Name()
		domain := r.loadDomainResolverInfo(name)
		domains = append(domains, domain)
	}

	return domains, nil
}

func (r *Resolve) loadDomainResolverInfo(name string) Domain {
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
		ws := regexp.MustCompile(`\s+`)
		parts := ws.Split(line, -1)
		if len(parts) != 2 {
			// TODO Panic
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
	return domain
}
