package main

import (
	"fmt"
	"strconv"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silverResult(lines []string) int {
	largestCount := 0
	currentSum := 0
	for _, l := range lines {
		if l == "" {
			if currentSum > largestCount {
				largestCount = currentSum
			}
			currentSum = 0
		} else {
			c, _ := strconv.Atoi(l)
			currentSum += c
		}
		if currentSum > largestCount {
			largestCount = currentSum
		}
	}
	return largestCount
}

func sum(lines []string) (s int) {
	for _, l := range lines {
		x, _ := strconv.Atoi(l)
		s = s + x
	}
	return s
}

func silver() {
	f := util.ReadFile("src/1/input.txt")
	count := silverResult(f)
	fmt.Println(count)
}

func gold() {
	// Print gold result
}
