package main

import (
	"fmt"
	"strconv"
	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
}

func IncrementalValues(input []string) int {
	incremented := 0
	var previous int
	for _, level := range input {
		cur, _ := strconv.Atoi(level)
		if previous != 0 && previous < cur {
			incremented++
		}
		previous = cur
	}
	return incremented
}

func silver() {
	lines := util.ReadFile("./src/1/input.txt")
	incremented := IncrementalValues(lines)
	fmt.Println(incremented)

}

