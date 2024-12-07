package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestLine1(t *testing.T) {
	tst.Equals(t, true, isSafe("7 6 4 2 1"))
	tst.Equals(t, true, isSafe("1 3 6 7 9"))
	tst.Equals(t, false, isSafe("1 2 7 8 9"))
	tst.Equals(t, false, isSafe("9 7 6 2 1"))
	tst.Equals(t, false, isSafe("1 3 2 4 5"))
	tst.Equals(t, false, isSafe("8 6 4 4 1"))
}
