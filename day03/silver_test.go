package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestParse(t *testing.T) {
	type testFixture struct {
		input  string
		parsed claim
	}
	t.Log("Parse")
	tests := []testFixture{
		{"#1 @ 520,746: 4x20", claim{520, 746, 4, 20, 1}},
		{"#2 @ 3,1: 4x4", claim{3, 1, 4, 4, 2}},
		{"#3 @ 5,5: 2x2", claim{5, 5, 2, 2, 3}},
	}
	for _, test := range tests {
		parsedData := parseLine(test.input)
		util.Equals(t, parsedData, test.parsed)
	}
}

func TestMaxWidthHeight(t *testing.T) {
	t.Log("Max width and height")
	type testFixture struct {
		input     []claim
		maxWidth  int
		maxHeight int
	}
	tests := []testFixture{
		{[]claim{claim{3, 1, 123, 231, 0}, claim{4, 4, 8, 3, 0}, claim{9, 1, 2, 6, 0}},
			126,
			232},
		{[]claim{claim{3, 1, 4, 4, 0}, claim{4, 4, 8, 3, 0}, claim{9, 1, 2, 6, 0}},
			12,
			7},
	}
	for _, test := range tests {
		maxWidth, maxHeight := getMaxes(test.input)
		util.Equals(t, test.maxHeight, maxHeight)
		util.Equals(t, test.maxWidth, maxWidth)
	}
}
func TestCovers(t *testing.T) {
	t.Log("Covers")
	type testFixture struct {
		input  claim
		x      int
		y      int
		covers bool
	}
	tests := []testFixture{
		{claim{3, 2, 5, 4, 0}, 0, 0, false},
		{claim{3, 2, 5, 4, 0}, 3, 2, true},
		{claim{3, 2, 5, 4, 0}, 7, 2, true},
		{claim{3, 2, 5, 4, 0}, 8, 2, false},
		{claim{3, 2, 5, 4, 0}, 7, 1, false},
		{claim{3, 2, 5, 4, 0}, 7, 6, false},
	}
	for _, test := range tests {
		covers := covers(test.input, test.x, test.y)
		t.Log("Covers", test)
		util.Equals(t, covers, test.covers)
	}
}

