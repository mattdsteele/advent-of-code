package main

import (
	"fmt"
	"strconv"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/3/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	sum := 0
	for _, i := range input {
		sum += silverJoltage(i)
	}
	return strconv.Itoa(sum)
}

func gold() {
	lines := util.ReadFile("src/3/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func silverJoltage(input string) int {
	// let's brute force this
	largest := 0
	for i, s := range input {
		substr := input[i+1:]
		for _, q := range substr {
			v := fmt.Sprintf("%s%s", string(s), string(q))
			candidate, _ := strconv.Atoi(v)
			if candidate > largest {
				largest = candidate
			}
		}
	}
	return largest
}
