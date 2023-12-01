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
	tst.Equals(t, "12", calibrationValue("1abc2"))
	tst.Equals(t, "38", calibrationValue("pqr3stu8vwx"))
}

func TestSilver(t *testing.T) {
	exampleInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	answer := silverCalculate(util.SliceAtLine(exampleInput))
	tst.Equals(t, "142", answer)
}
