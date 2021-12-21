package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
}

var initialState = "3,4,3,1,2"
var game = parse(initialState)

func TestInitialLoad(t *testing.T) {
	tst.Equals(t, 5, game.fishCount())
}

func TestOneGame(t *testing.T) {
	game.tick()
	tst.Equals(t, 1, game.rounds)
	tst.SlicesEqual([]int{2, 3, 2, 0, 1}, game.fishes)
}

func TestTwoRounds(t *testing.T) {
	game.tick()
	game.tick()
	tst.Equals(t, 2, game.rounds)
	tst.SlicesEqual([]int{1, 2, 1, 6, 0, 8}, game.fishes)
}
func Test80Days(t *testing.T) {
	game.playRounds(80)
	tst.Equals(t, 5934, game.fishCount())
}
func Test256(t *testing.T) {
	game.playRounds(256)
	tst.Equals(t, 26984457539, game.fishCount())
}
