package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestSilver(t *testing.T) {
	exampleInput := `125 17`
	nextLine := silverIterate(exampleInput)
	tst.Equals(t, nextLine, "253000 1 7")

	tst.Equals(t, silverIterate(silverIterate(exampleInput)), "253 0 2024 14168")

	tst.Equals(t, "55312", silverCalculate([]string{exampleInput}))
}
