package testing

import (
	"reflect"
	"strings"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

type Spec struct {
	t *testing.T
}

type SpecResult struct {
	t       *testing.T
	actuals []interface{}
}

// Create a spec test.
func SpecTest(t *testing.T) *Spec {
	return &Spec{t: t}
}

// Set the actuals.
func (s *Spec) Expect(actuals ...interface{}) (sr *SpecResult) {
	return &SpecResult{t: s.t, actuals: actuals}
}

// Check if the expectation equals the real value.
func (sr *SpecResult) ToEqual(expectations ...interface{}) {
	for index, expectation := range expectations {
		if !reflect.DeepEqual(expectation, sr.actuals[index]) {
			sr.t.Errorf("\t%s\t\"%+v\" does not equal \"%+v\"", ballotX, sr.actuals[index], expectation)
		}
	}
}

// Check if the a partial string is included.
func (sr *SpecResult) ToContain(expectations ...string) {
	for index, expectation := range expectations {
		switch v := sr.actuals[index].(type) {
		case string:
			if !strings.Contains(v, expectation) {
				sr.t.Errorf("\t%s\t\"%+v\" does not contain \"%+v\"", ballotX, v, expectation)
			}
			break
		default:
			sr.t.Errorf("\t%s\t\"%+v\" is not a string", ballotX, v)
		}
	}
}

// Check if the expectation does not equal the real value.
func (sr *SpecResult) ToNotEqual(expectations ...interface{}) {
	for index, expectation := range expectations {
		if !reflect.DeepEqual(expectation, sr.actuals[index]) {
			sr.t.Errorf("\t%s\t\"%+v\" does equal \"%+v\"", ballotX, sr.actuals[index], expectation)
		}
	}
}
