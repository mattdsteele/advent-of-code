package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	tst.Equals(t, 7, silverCalculate(exampleInput))
	tst.Equals(t, 5, silverCalculate("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	tst.Equals(t, 11, silverCalculate("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))

	tst.Equals(t, 19, goldCalculate("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	tst.Equals(t, 26, goldCalculate("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
