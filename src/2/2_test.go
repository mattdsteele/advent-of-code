package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestSilver(t *testing.T) {
	tst.Equals(t, true, silverSafe("7 6 4 2 1"))
	tst.Equals(t, true, silverSafe("1 3 6 7 9"))
	tst.Equals(t, false, silverSafe("1 2 7 8 9"))
	tst.Equals(t, false, silverSafe("9 7 6 2 1"))
	tst.Equals(t, false, silverSafe("1 3 2 4 5"))
	tst.Equals(t, false, silverSafe("8 6 4 4 1"))
}

func TestGold(t *testing.T) {
	tst.Equals(t, true, goldSafe("7 6 4 2 1"))
	tst.Equals(t, true, goldSafe("1 3 6 7 9"))
	tst.Equals(t, false, goldSafe("1 2 7 8 9"))
	tst.Equals(t, false, goldSafe("9 7 6 2 1"))
	tst.Equals(t, true, goldSafe("1 3 2 4 5"))
	tst.Equals(t, true, goldSafe("8 6 4 4 1"))
}
