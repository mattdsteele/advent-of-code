package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	input := `1
2
3
5
`
	lines := util.SliceAtLine(input)
	count := sum(lines)
	tst.Equals(t, 11, count)
}

func TestSilver(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	lines := util.SliceAtLine(input)
	tst.Equals(t, 24000, silverResult(lines))
}
func TestGold(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	lines := util.SliceAtLine(input)
	tst.Equals(t, 45000, goldResult(lines))

}
