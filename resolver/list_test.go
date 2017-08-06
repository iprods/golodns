package resolver

import (
	"os"
	"testing"

	specs "github.com/iprods/golodns/testing"
)

func Test_it_handles_a_not_existing_directory(t *testing.T) {
	spec := specs.SpecTest(t)
	r := Resolve{
		Path: "/tmp/resolver",
	}
	_, err := r.List()
	spec.Expect(err.Error()).ToEqual("open /tmp/resolver: no such file or directory")
}

func Test_it_loads_an_existing_entry_with_only_a_nameserver(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, os.FileMode(0777))
	defer os.RemoveAll("/tmp/golodns")
	f, _ := os.Create(path + "/dns")
	f.WriteString("# Managed by golodns\n")
	f.WriteString("nameserver 192.168.1.1\n")
	f.Close()
	r := Resolve{
		Path: path,
	}
	domains, _ := r.List()
	spec.Expect(len(domains)).ToEqual(1)
	expectedDomain := Domain{
		Name:      "dns",
		IpAddress: "192.168.1.1",
		Port:      "53",
		Managed:   true,
	}
	spec.Expect(domains[0]).ToEqual(expectedDomain)
}

func Test_it_loads_an_existing_entry_with_a_nameserver_and_port(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, os.FileMode(0777))
	defer os.RemoveAll("/tmp/golodns")
	f, _ := os.Create(path + "/dns")
	f.WriteString("# Managed by golodns\n")
	f.WriteString("nameserver 192.168.1.1\n")
	f.WriteString("port       8053\n")
	f.Close()
	r := Resolve{
		Path: path,
	}
	domains, _ := r.List()
	spec.Expect(len(domains)).ToEqual(1)
	expectedDomain := Domain{
		Name:      "dns",
		IpAddress: "192.168.1.1",
		Port:      "8053",
		Managed:   true,
	}
	spec.Expect(domains[0]).ToEqual(expectedDomain)
}

func Test_it_loads_an_existing_entry_that_is_not_managed(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, os.FileMode(0777))
	defer os.RemoveAll("/tmp/golodns")
	f, _ := os.Create(path + "/custom")
	f.WriteString("nameserver 127.0.0.1\n")
	f.Close()
	r := Resolve{
		Path: path,
	}
	domains, _ := r.List()
	spec.Expect(len(domains)).ToEqual(1)
	expectedDomain := Domain{
		Name:      "custom",
		IpAddress: "127.0.0.1",
		Port:      "53",
		Managed:   false,
	}
	spec.Expect(domains[0]).ToEqual(expectedDomain)
}

func Test_it_errors_out_if_a_malformed_line_is_encountered(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, os.FileMode(0777))
	defer os.RemoveAll("/tmp/golodns")
	f, _ := os.Create(path + "/dns")
	f.WriteString("nameserver 127.0.0.1 blabla\n")
	f.Close()
	r := Resolve{
		Path: path,
	}
	_, err := r.List()
	spec.Expect(err.Error()).ToEqual("Unexpected line. Found 3 parts.")
}

func Test_it_ignores_comment_lines(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, os.FileMode(0777))
	defer os.RemoveAll("/tmp/golodns")
	f, _ := os.Create(path + "/dns")
	f.WriteString("# Custom comment\n")
	f.WriteString("nameserver 127.0.0.1\n")
	f.WriteString("# Another custom comment\n")
	f.Close()
	r := Resolve{
		Path: path,
	}
	domains, _ := r.List()
	spec.Expect(len(domains)).ToEqual(1)
	expectedDomain := Domain{
		Name:      "dns",
		IpAddress: "127.0.0.1",
		Port:      "53",
		Managed:   false,
	}
	spec.Expect(domains[0]).ToEqual(expectedDomain)
}

func Test_it_ignores_comments_in_the_same_line(t *testing.T) {
	spec := specs.SpecTest(t)
	path := "/tmp/golodns/resolver"
	os.MkdirAll(path, os.FileMode(0777))
	defer os.RemoveAll("/tmp/golodns")
	f, _ := os.Create(path + "/dns")
	f.WriteString("nameserver 127.0.0.1 # Custom comment\n")
	f.Close()
	r := Resolve{
		Path: path,
	}
	domains, _ := r.List()
	spec.Expect(len(domains)).ToEqual(1)
	expectedDomain := Domain{
		Name:      "dns",
		IpAddress: "127.0.0.1",
		Port:      "53",
		Managed:   false,
	}
	spec.Expect(domains[0]).ToEqual(expectedDomain)
}
