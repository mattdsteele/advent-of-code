package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
}

func TestBreakupLines(t *testing.T) {
	data := `1
2
3
4
5`
	lines := util.SliceAtLine(data)
	vals := Chunked(lines, 3)
	tst.Equals(t, len(vals), 3)
	tst.SlicesEqual([]int{1, 2, 3}, lines[0])
	tst.SlicesEqual([]int{2, 3, 4}, lines[1])
}

func TestGoldExample(t *testing.T) {
	data := `199
200
208
210
200
207
240
269
260
263`
	gr := goldResult(util.SliceAtLine(data))
	tst.Equals(t, 5, gr)
}

func TestUpDown(t *testing.T) {
	data := `1
2
1
4
2`
	lines := util.SliceAtLine(data)
	val := IncrementalValues(lines)
	tst.Equals(t, val, 2)

}
