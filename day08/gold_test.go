package main
import (
	"strconv"
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestNoChildrenGold(t *testing.T) {
	fixtures := [][]string{
		[]string{"0 1 99", "99"},
		[]string{"0 2 3 33", "36"},
		[]string{"0 3 2 33 1", "36"},
		[]string{"0 4 2 33 1 0", "36"},
	}
	for _, fixture := range fixtures {
		sum := strconv.Itoa(goldValue(fixture[0], 0))
		util.Equals(t, fixture[1], sum)
	}
}
func TestSingleChildGold(t *testing.T) {
	fixtures := [][]string{
		[]string{"1 1 0 1 99 1", "0", "99"},
		[]string{"1 2 0 1 99 2 3", "0", "0"},
		[]string{"1 2 0 1 99 2 0", "0", "0"},
		[]string{"1 2 0 1 99 2 1", "0", "99"},
	}
	for _, fixture := range fixtures {
		idx, _ := strconv.Atoi(fixture[1])
		sum := strconv.Itoa(goldValue(fixture[0], idx))
		util.Equals(t, fixture[2], sum)
	}
}

func TestExampleGold(t *testing.T) {
	sum := goldValue("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", 0)
	util.Equals(t, sum, 66)
}

