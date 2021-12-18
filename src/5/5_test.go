package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sample string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

var in []string = util.SliceAtLine(sample)

func TestParse(t *testing.T) {
	//tst.Equals(t, 5, overlappingLines(in))
}
func TestParseLine(t *testing.T) {
	line := parseLine(in[0])
	tst.Equals(t, 0, line.from.x)
	tst.Equals(t, 9, line.to.y)
}

func TestRange(t *testing.T) {
	r := toRange(parseLine(in[0]), silverLineStrat)
	tst.Equals(t, 6, len(r))
	tst.Equals(t, r[1], &point{1, 9})

	r = toRange(parseLine(in[2]), silverLineStrat)
	tst.Equals(t, 7, len(r))
	tst.Equals(t, r[1], &point{8, 4})
}

func TestPoints(t *testing.T) {
	g := setup(in)
	tst.Equals(t, 5, g.overlapping())
}

func TestGold(t *testing.T) {
	g := setup(in)
	tst.Equals(t, 12, g.goldOverlapping())
}

func TestGoldRange(t *testing.T) {
	r := toRange(parseLine("1,1 -> 3,3"), goldLineStrat)
	tst.Equals(t, 3, len(r))
	tst.Equals(t, r[1], &point{2, 2})

	r = toRange(parseLine("9,7 -> 7,9"), goldLineStrat)
	tst.Equals(t, 3, len(r))
	tst.Equals(t, r[1], &point{8, 8})
}
