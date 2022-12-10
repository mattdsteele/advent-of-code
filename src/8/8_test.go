package main

import (
	"fmt"
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `30373
25512
65332
33549
35390`
	lines := util.SliceAtLine(exampleInput)
	board := parseBoard(lines)
	tst.Equals(t, 5, board.width)
	tst.Equals(t, 5, board.height)
	t0 := board.get(0, 0)
	fmt.Println(t0)
	tst.Equals(t, 3, t0.height)
	tst.Equals(t, 1, board.get(1, 3).height)

	tst.Equals(t, true, board.visible(0, 0))
	tst.Equals(t, true, board.visible(1, 0))
	tst.Equals(t, true, board.visible(0, 1))

	tst.Equals(t, false, board.visible(1, 3))
	tst.Equals(t, true, board.visible(1, 2))

	tst.Equals(t, 21, silverCalculate(lines))
}
