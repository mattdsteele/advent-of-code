package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var exampleInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestSilver(t *testing.T) {
	tst.Equals(t, 2, 2)
	lines := util.SliceAtLine(exampleInput)
	data := parse(lines)
	exampleRange := []int{3, 5}
	tst.Equals(t, true, inRange(exampleRange, 3))
	tst.Equals(t, false, inRange(exampleRange, 6))
	tst.Equals(t, true, data.inRange(3))
	tst.Equals(t, false, data.inRange(6))
	tst.Equals(t, 3, data.silver())
}

func TestGold(t *testing.T) {
	lines := util.SliceAtLine(exampleInput)
	data := parse(lines)
	tst.Equals(t, 14, data.gold())
}
