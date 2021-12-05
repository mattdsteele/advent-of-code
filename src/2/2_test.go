package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	sampleData := util.SliceAtLine(exampleData())
	tst.Equals(t, 150, silverData(sampleData))
}

func TestParseLine(t *testing.T) {
	dir, amt := parseLine("forward 5")
	tst.Equals(t, "forward", dir)
	tst.Equals(t, 5, amt)
}

func exampleData() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}
