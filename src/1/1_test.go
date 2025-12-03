package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var exampleInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestSilver(t *testing.T) {
	lines := util.SliceAtLine(exampleInput)
	answer := silverCalculate(lines)
	tst.Equals(t, answer, "3")
}

func TestGold(t *testing.T) {
	lines := util.SliceAtLine(exampleInput)
	answer := goldCalculate(lines)
	tst.Equals(t, answer, "6")
}

func TestRange(t *testing.T) {
	tst.Equals(t, -30, delta("L30"))
	tst.Equals(t, 55, delta("R55"))
}

func TestInc(t *testing.T) {
	tst.Equals(t, 1, 105/100)
	tst.Equals(t, 2, 210/100)
	tst.Equals(t, -3, -305/100)
}
