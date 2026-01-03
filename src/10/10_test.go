package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var input = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`
var lines = util.SliceAtLine(input)

func TestParse(t *testing.T) {
	m := parse(lines[0])
	state := []bool{false, false, false, false}
	newState := push(state, []int{2, 3, 0})
	tst.Equals(t, false, state[2])
	tst.Equals(t, true, newState[2])
	tst.Equals(t, "....", str(state))
	tst.Equals(t, "#.##", str(newState))

	tst.Equals(t, false, matches(state, newState))
	tst.Equals(t, true, matches(state, push(state, []int{})))

	s2 := push(newState, []int{0, 1})
	tst.Equals(t, ".###", str(s2))
	tst.Equals(t, 2, m.silver())

	tst.Equals(t, 3, parse(lines[1]).silver())
	tst.Equals(t, 2, parse(lines[2]).silver())

	tst.Equals(t, 7, game(lines).silver())
}
