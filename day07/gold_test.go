package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestStepTime(t *testing.T) {
	util.Equals(t, 61, stepTime("A", 60))
	util.Equals(t, 86, stepTime("Z", 60))
	util.Equals(t, 1, stepTime("A", 0))
	util.Equals(t, 26, stepTime("Z", 0))
}

func TestRunner(t *testing.T) {
	totalTime := timeWithHelpers(reqs(), 0, 2)
	util.Equals(t, 15, totalTime)
}

