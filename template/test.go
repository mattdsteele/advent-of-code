package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `exampl`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, lines, "1234")
}
