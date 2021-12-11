package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var input string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

var lines []string = util.SliceAtLine(input)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
}

func TestSample(t *testing.T) {
	tst.Equals(t, 198, silverResult(lines))
}

func TestOxygenRating(t *testing.T) {
	tst.Equals(t, 23, oxygenRating(lines))
}

func TestOxygenStep(t *testing.T) {
	nextStep := []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001"}
	tst.Equals(t, nextStep, oxygenStep(lines, 0))
	next := []string{"10110", "10111", "10101", "10000"}
	tst.Equals(t, next, oxygenStep(nextStep, 1))
	next3 := []string{"10110", "10111", "10101"}
	tst.Equals(t, next3, oxygenStep(next, 2))

}
