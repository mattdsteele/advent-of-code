package main

import (
	"fmt"
	"regexp"
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
	total := 0
	for _, l := range input {
		total += silverSolve(l)
	}
	return strconv.Itoa(total)
}

func gold() {
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func silverSolve(input string) int {
	re := regexp.MustCompile(`mul\(([\d]+),([\d]+)\)`)
	matches := re.FindAllStringSubmatch(input, -1) // Find all matches

	total := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += (x * y)
	}
	return total
}
