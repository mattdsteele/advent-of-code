package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestAllBlank(t *testing.T) {
	t.Log("Covers")
	type testFixture struct {
		inputs    []claim
		exclusive bool
	}
	tests := []testFixture{
		{[]claim{claim{3, 2, 5, 4, 0}}, true},
		{[]claim{claim{3, 2, 5, 4, 0}, claim{3, 2, 6, 4, 0}}, false},
	}
	for _, test := range tests {
		exclusive := isExclusive(test.inputs, test.inputs[0])
		t.Log("Is exclusive", test.inputs)
		util.Equals(t, exclusive, test.exclusive)
	}
}

