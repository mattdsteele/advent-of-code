package main

import (
	"sort"
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestDistance(t *testing.T) {
	tst.Equals(t, 3, distance(2, 5))
	tst.Equals(t, 3, distance(5, 2))
}

func TestParseLines(t *testing.T) {
	first, second := getEntries("80784   47731")
	tst.Equals(t, first, 80784)
	tst.Equals(t, second, 47731)
}

func TestSort(t *testing.T) {
	unsorted := []int{1, 2, 4, 3}
	sort.Ints(unsorted)
	tst.Equals(t, []int{1, 2, 3, 4}, unsorted)
}
