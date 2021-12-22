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
	tst.Equals(t, 41, g.fuelCost(1, fuel))
	tst.Equals(t, 71, g.fuelCost(10, fuel))
}

func TestBestPosition(t *testing.T) {
	g := parse(testInput)
	tst.Equals(t, 37, g.bestFuelCost())
}

func TestGoldFuelStrat(t *testing.T) {
	tst.Equals(t, 66, goldFuelStrat(16, 5))
	tst.Equals(t, 66, goldFuelStrat(5, 16))
	tst.Equals(t, 6, goldFuelStrat(2, 5))
	tst.Equals(t, 45, goldFuelStrat(14, 5))
}

func TestGold(t *testing.T) {
	g := parse(testInput)
	tst.Equals(t, 168, g.bestGoldFuelCost())
}
