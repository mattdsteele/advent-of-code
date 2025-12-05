package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestSilver(t *testing.T) {
	exampleInput := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, silverCalculate(lines), "1227775554")
}
func TestGold(t *testing.T) {
	exampleInput := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, goldCalculate(lines), "4174379265")
}

func TestValid(t *testing.T) {
	tst.Equals(t, valid(11), false)
	tst.Equals(t, valid(12), true)
	tst.Equals(t, valid(22), false)
	tst.Equals(t, valid(1188511885), false)
	tst.Equals(t, valid(1698522), true)
}
func TestGoldValid(t *testing.T) {
	tst.Equals(t, goldValid(11), false)
	tst.Equals(t, goldValid(12), true)
	tst.Equals(t, goldValid(22), false)
	tst.Equals(t, goldValid(99), false)
	tst.Equals(t, goldValid(100), true)
	tst.Equals(t, goldValid(111), false)
	tst.Equals(t, goldValid(1010), false)
}
