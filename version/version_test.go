package version

import (
	"regexp"
	"testing"

	specs "github.com/iprods/golodns/testing"
)

func Test_it_prints_the_human_readable_version_number(t *testing.T) {
	spec := specs.SpecTest(t)
	var format, _ = regexp.Compile(`^\d*\.\d*.\d*(?:-dev)?$`)
	spec.Expect(format.Match([]byte(HumanReadableVersion()))).ToEqual(true)
}
