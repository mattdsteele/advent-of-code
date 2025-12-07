package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var exampleInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestSilver(t *testing.T) {
	lines := util.SliceAtLine(exampleInput)
	g := parse(lines)
	total := g.surrounding(0, 2)
	tst.Equals(t, 3, total)

	silver := g.silver()
	tst.Equals(t, 13, silver)
}

func TestGold(t *testing.T) {
	lines := util.SliceAtLine(exampleInput)
	g := parse(lines)
	tst.Equals(t, 43, g.gold())
}
