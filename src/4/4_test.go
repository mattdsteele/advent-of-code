package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sample = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestFindForward(t *testing.T) {
	s := Search{}
	s.fields = util.SliceAtLine(`MMMSXXMASM
MSAMXMSMSA`)
	tst.Equals(t, 0, s.check(0, 0))
	tst.Equals(t, 1, s.check(0, 5))
	tst.Equals(t, 1, s.check(1, 4))
}

func TestSilverCalc(t *testing.T) {
	tst.Equals(t, "18", silverCalculate(util.SliceAtLine(sample)))
}
func TestFindAll(t *testing.T) {
	s := Search{}
	s.fields = util.SliceAtLine(sample)
	tst.Equals(t, 18, s.silver())
}

func TestGo(t *testing.T) {
	s := Search{}
	s.fields = util.SliceAtLine(sample)
	tst.Equals(t, true, s.goldCheck(1, 2))
	tst.Equals(t, 9, s.gold())
}
