package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sampleData = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestParse(t *testing.T) {
	lines := util.SliceAtLine(sampleData)
	p := parse(lines)
	tst.Equals(t, 50, p.silver())
}
