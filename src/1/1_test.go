package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	lines := util.SliceAtLine(exampleInput)
	tst.Equals(t, lines[0], "1abc2")
}

func TestDigits(t *testing.T) {
	tst.Equals(t, 12, silverCalibration("1abc2"))
	tst.Equals(t, 38, silverCalibration("pqr3stu8vwx"))
}

func TestSilver(t *testing.T) {
	exampleInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	answer := silverCalculate(util.SliceAtLine(exampleInput))
	tst.Equals(t, "142", answer)
}

func TestGold(t *testing.T) {
	exampleInput := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	answer := goldCalculate(util.SliceAtLine((exampleInput)))
	tst.Equals(t, "281", answer)
}

func TestGoldLine(t *testing.T) {
	tst.Equals(t, 29, goldCalibration("two1nine"))
	tst.Equals(t, 83, goldCalibration("eightwothree"))
}
