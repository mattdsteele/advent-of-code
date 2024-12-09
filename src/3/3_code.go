package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
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

func goldSolve(input string, total int) int {
	// calculate until find don't, then go again
	endIdx := strings.Index(input, `don't()`)
	if endIdx == -1 {
		total += silverSolve(input)
		return total
	}

	sec := input[:endIdx]
	total += silverSolve(sec)
	// find where to start again
	rem := input[endIdx:]
	startIdx := strings.Index(rem, `do()`)
	if startIdx == -1 {
		return total
	}
	return goldSolve(rem[startIdx:], total)
}

func gold() {
	lines := util.ReadFile("src/3/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	total := 0
	fullLines := ""
	for _, l := range lines {
		fullLines += l
	}
	total += goldSolve(fullLines, 0)
	return strconv.Itoa(total)
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
