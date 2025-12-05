package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestSilver(t *testing.T) {
	exampleInput := `987654321111111
811111111111119
234234234234278
818181911112111`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, silverJoltage(lines[0]), 98)
	tst.Equals(t, silverJoltage("811111111111119"), 89)
	tst.Equals(t, silverCalculate(lines), "357")
}
