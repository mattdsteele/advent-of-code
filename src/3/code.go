package main

import (
	"fmt"
	"strconv"
	"strings"

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

func priority(s string) int {
	asciiVal := fmt.Sprintf("%d", s[0])
	i, _ := strconv.Atoi(asciiVal)
	if i > 96 {
		return i - 96
	}
	return i - 38
}

func sharedItem(sacks []string) string {
	for _, s := range sacks[0] {
		if strings.ContainsRune(sacks[1], s) {
			return string(s)
		}
	}
	panic("could not find item")
}

func getRucksacks(exampleInput string) []string {
	inputSize := len(exampleInput)
	firstVal := exampleInput[0 : inputSize/2]
	secondVal := exampleInput[inputSize/2:]
	return []string{firstVal, secondVal}

}

func silverCalculate(input []string) (total int) {
	for _, i := range input {
		total += priority(sharedItem(getRucksacks(i)))
	}
	return total
}

func gold() {
	lines := util.ReadFile("src/3/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}
