package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
}

var initialState = "3,4,3,1,2"

func TestInitialLoad(t *testing.T) {
	var game = parse(initialState)
	tst.Equals(t, 5, game.fishCount())
}

func TestOneGame(t *testing.T) {
	var game = parse(initialState)
	game.tick()
	tst.Equals(t, 1, game.rounds)
	tst.Equals(t, 5, game.fishCount())
}

func TestTwoRounds(t *testing.T) {
	var game = parse(initialState)
	game.tick()
	game.tick()
	tst.Equals(t, 2, game.rounds)
	tst.Equals(t, 6, game.fishCount())
}
func Test10(t *testing.T) {
	var game = parse(initialState)
	game.playRounds(10)
	tst.Equals(t, 12, game.fishCount())
}
func Test18(t *testing.T) {
	var game = parse(initialState)
	game.playRounds(18)
	tst.Equals(t, 26, game.fishCount())
}
func Test80Days(t *testing.T) {
	var game = parse(initialState)
	game.playRounds(80)
	tst.Equals(t, 5934, game.fishCount())
}
func Test256(t *testing.T) {
	var game = parse(initialState)
	game.playRounds(256)
	tst.Equals(t, 26984457539, game.fishCount())
}
