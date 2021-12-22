package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

var testInput = "16,1,2,0,4,2,7,1,2,14"

func TestParse(t *testing.T) {
	g := parse(testInput)
	tst.Equals(t, 10, len(g.crabPositions))
}

func TestMinMaxValues(t *testing.T) {
	g := parse(testInput)
	tst.Equals(t, 16, g.max())
	tst.Equals(t, 0, g.min())
}

func TestFuelCost(t *testing.T) {
	g := parse(testInput)
	tst.Equals(t, 41, g.fuelCost(1))
	tst.Equals(t, 71, g.fuelCost(10))
}

func TestBestPosition(t *testing.T) {
	g := parse(testInput)
	tst.Equals(t, 37, g.bestFuelCost())

}
