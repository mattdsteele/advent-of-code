package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sample = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestParseHeader(t *testing.T) {
	game := parse(util.SliceAtLine(sample))
	tst.Assert(t, game != nil, "not nil")
	tst.Equals(t, 3, len(game.columns))

	tst.Equals(t, 2, len(game.columns[0].items))
	tst.Equals(t, 1, len(game.columns[2].items))

	tst.Assert(t, tst.SlicesEqual([]string{"N", "Z"}, game.columns[0].items), "wrong messages")
	tst.Assert(t, tst.SlicesEqual([]string{"D", "C", "M"}, game.columns[1].items), "wrong messages")
}

func TestBody(t *testing.T) {
	step := parseStep("move 1 from 2 to 1")
	tst.Equals(t, 1, step.count)
	tst.Equals(t, 2, step.from)
	tst.Equals(t, 1, step.to)
}

func TestSilverGame(t *testing.T) {
	game := parse(util.SliceAtLine(sample))
	tst.Equals(t, "CMZ", game.playSilver())
}

func TestGoldGame(t *testing.T) {
	game := parse(util.SliceAtLine(sample))
	tst.Equals(t, "MCD", game.playGold())
}
