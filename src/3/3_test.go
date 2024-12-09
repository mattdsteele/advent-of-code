package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestExample(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	tst.Equals(t, 161, silverSolve(input))
}

func TestGold(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	tst.Equals(t, 48, goldSolve(input, 0))
}
