package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
	util "github.com/mattdsteele/advent-of-code"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
}

func TestIncremental(t *testing.T) {
	data := `1
2
3
4
5`
	lines := util.SliceAtLine(data)
	val := IncrementalValues(lines)
	tst.Equals(t, val, 4)


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