package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestSilver(t *testing.T) {
	input := `A Y
B X
C Z`
	lines := util.SliceAtLine(input)
	tst.Equals(t, 15, silverResult(lines))
}

func TestParse(t *testing.T) {
	line := "A Y"
	tst.Equals(t, 8, silverPoints(line))

}
