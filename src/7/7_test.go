package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sampleInput = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := sampleInput
	lines := util.SliceAtLine(exampleInput)
	game := parse(lines)

	vals := make(map[int]bool)
	vals[7] = true
	game.beams = vals
	splits := game.silverTick(lines[2])
	tst.Equals(t, 1, splits)
	tst.Equals(t, vals[6], true)
	tst.Equals(t, vals[7], false)
	tst.Equals(t, vals[8], true)

	game = parse(lines)
	splits = game.silverTick(lines[0])
	tst.Equals(t, 0, splits)
	tst.Equals(t, game.beams[6], false)
	tst.Equals(t, game.beams[7], true)
	tst.Equals(t, game.beams[8], false)

	tst.Equals(t, 21, game.silver())
}

func TestGold(t *testing.T) {
	lines := util.SliceAtLine(sampleInput)
	game := parse(lines)
	tst.Equals(t, 40, game.gold())
}
