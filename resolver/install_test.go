package resolver

import (
	"os"
	"testing"

	specs "github.com/iprods/golodns/testing"
	"io/ioutil"
)

func Test_it_creates_a_new_domain_resolver_with_the_needed_parent_directory(t *testing.T) {
	spec := specs.SpecTest(t)
	os.Mkdir("/tmp/golodns", 0777)
	defer os.RemoveAll("/tmp/golodns")
	_, err := os.Open("/tmp/golodns/resolver")
	spec.Expect(err.Error()).ToEqual("open /tmp/golodns/resolver: no such file or directory")
	r := Resolve{
		Path: "/tmp/golodns/resolver",
	}
	domainName := "some"
	addr := "127.0.0.1"
	port := "5300"
	domain, _ := r.Install(domainName, addr, port)
	contents, _ := ioutil.ReadFile("/tmp/golodns/resolver/some")
	spec.Expect(string(contents)).ToEqual("# Managed by golodns\nnameserver 127.0.0.1\nport 5300\n")
	spec.Expect(domain.Name).ToEqual("some")
}

func Test_it_skips_an_already_existing_domain_resolver(t *testing.T) {
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
	addr := "127.0.0.1"
	port := "5300"
	_, err = r.Install(domainName, addr, port)
	spec.Expect(err.Error()).ToEqual("Domain tld is already being resolved.")
}
