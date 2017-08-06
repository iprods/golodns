package resolver

import (
	"os"
	"testing"

	specs "github.com/iprods/golodns/testing"
)

func Test_it_tells_tried_removal_of_a_not_existing_domain_resolver(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, 0777)
	defer os.RemoveAll("/tmp/golodns")
	domainName := "tld"
	r := Resolve{
		Path: path,
	}
	_, err := r.Uninstall(domainName)
	spec.Expect(err.Error()).ToEqual("Domain resolver for tld does not exist.")
}

func Test_it_tells_only_managed_existing_domain_resolvers_can_be_removed(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, 0777)
	defer os.RemoveAll("/tmp/golodns")
	domainName := "tld"
	resFile := path + "/" + domainName
	f, _ := os.Create(resFile)
	f.WriteString("nameserver 192.168.1.1\n")
	f, err := os.Open(resFile)
	defer f.Close()
	spec.Expect(err).ToEqual(nil)
	r := Resolve{
		Path: path,
	}
	_, err = r.Uninstall(domainName)
	spec.Expect(err.Error()).ToEqual("Domain tld is not managed.")
}

func Test_it_removes_a_domain_resolver_when_it_is_managed(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, 0777)
	defer os.RemoveAll("/tmp/golodns")
	domainName := "tld"
	resFile := path + "/" + domainName
	f, _ := os.Create(resFile)
	f.WriteString("# Managed by golodns\n")
	f.WriteString("nameserver 192.168.1.1\n")
	f, err := os.Open(resFile)
	defer f.Close()
	spec.Expect(err).ToEqual(nil)
	r := Resolve{
		Path: path,
	}
	domain, _ := r.Uninstall(domainName)
	spec.Expect(domain.Name).ToEqual(domainName)
	spec.Expect(domain.Managed).ToEqual(false)
	_, err = os.Open(resFile)
	defer f.Close()
	spec.Expect(err.Error()).ToEqual("open /tmp/golodns/resolver/tld: no such file or directory")
}
