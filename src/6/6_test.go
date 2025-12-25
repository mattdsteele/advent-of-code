package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestSilver(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	lines := util.SliceAtLine(exampleInput)
	data := parse(lines)
	tst.Equals(t, 4277556, data.silver())
}

func TestGold(t *testing.T) {
	exampleInput := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	lines := util.SliceAtLine(exampleInput)
	data := goldParse(lines)
	tst.Equals(t, 3263827, data.gold())
}
