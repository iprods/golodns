package resolver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Install the file in the resolver directory
func (r *Resolve) Install(name string, addr string, port string) (Domain, error) {
	domain := Domain{
		Name:      name,
		IpAddress: addr,
		Port:      port,
		Managed:   true,
	}
	path := r.Path
	files, err := ioutil.ReadDir(path)
	// The outermost directory is missing so try to create it; Absent by default
	if err != nil {
		mkErr := os.Mkdir(path, 0755)
		if mkErr != nil {
			return Domain{}, mkErr
		}
	}
	// Check if the domain is already managed
	for _, file := range files {
		if name == file.Name() {
			return Domain{}, errors.New(fmt.Sprintf("Domain %s is already being resolved.", name))
		}
	}
	f, err := os.Create(path + "/" + name)
	if err != nil {
		return Domain{}, err
	}
	_, err = f.WriteString(fmt.Sprintf("# Managed by golodns\nnameserver %s\nport %s\n", addr, port))
	if err != nil {
		return Domain{}, err
	}
	f.Close()
	return domain, nil
}
