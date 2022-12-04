package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `2-4,6-8
2-8,3-7
6-6,4-6`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, false, enclosed(lines[0]))
	tst.Equals(t, true, enclosed(lines[1]))
	tst.Equals(t, true, enclosed(lines[2]))
}

func TestSilver(t *testing.T) {
	exampleInput := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, 2, silverCalculate(lines))

}
