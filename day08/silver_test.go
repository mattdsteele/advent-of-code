package main
import (
	"strconv"
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestOneNode(t *testing.T) {
	fixtures := [][]string{
		[]string{"0 1 99", "99"},
		[]string{"0 2 3 33", "36"},
		[]string{"0 3 2 33 1", "36"},
		[]string{"0 4 2 33 1 0", "36"},
	}
	for _, fixture := range fixtures {
		sum := strconv.Itoa(checksum(fixture[0], 0))
		util.Equals(t, fixture[1], sum)
	}
}
func TestSubLeafNode(t *testing.T) {
	fixtures := [][]string{
		[]string{"1 1 0 1 99 2", "2", "99"},
		[]string{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "9", "99"},
		[]string{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "2", "33"},
	}
	for _, fixture := range fixtures {
		idx, _ := strconv.Atoi(fixture[1])
		sum := strconv.Itoa(checksum(fixture[0], idx))
		util.Equals(t, fixture[2], sum)
	}
}
func TestSingleChild(t *testing.T) {
	fixtures := [][]string{
		[]string{"1 1 0 1 99 2", "0", "101"},
		[]string{"1 1 1 1 0 1 10 99 2", "0", "111"},
	}
	for _, fixture := range fixtures {
		idx, _ := strconv.Atoi(fixture[1])
		sum := strconv.Itoa(checksum(fixture[0], idx))
		util.Equals(t, fixture[2], sum)
	}
}

func TestMultipleChildren(t *testing.T) {
	fixtures := [][]string{
		[]string{"2 1 0 1 99 0 1 55 2", "0", "156"},
		[]string{"3 1 0 1 99 0 1 55 0 1 22 2", "0", "178"},
		[]string{"3 1 1 1 0 1 3 99 0 1 55 0 1 22 2", "0", "181"},
	}
	for _, fixture := range fixtures {
		idx, _ := strconv.Atoi(fixture[1])
		sum := strconv.Itoa(checksum(fixture[0], idx))
		util.Equals(t, fixture[2], sum)
	}
}

func TestExample(t *testing.T) {
	sum := checksum("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", 0)
	util.Equals(t, sum, 138)
}

