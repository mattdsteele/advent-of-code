package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	lines := util.SliceAtLine(exampleInput)
	answer := silverCalculate(lines)
	tst.Equals(t, answer, "3")
}

func TestRange(t *testing.T) {
	tst.Equals(t, -30, delta("L30"))
	tst.Equals(t, 55, delta("R55"))
}
