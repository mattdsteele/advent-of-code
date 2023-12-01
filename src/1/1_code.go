package main

import (
	"fmt"
	"strconv"
	"unicode"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	total := 0
	for _, line := range input {
		val := calibrationValue((line))
		num, _ := strconv.Atoi(val)
		total += num
	}
	return fmt.Sprint(total)
}

func gold() {
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func calibrationValue(line string) string {
	var firstValue rune
	var lastValue rune
	for _, bar := range line {
		if unicode.IsDigit((bar)) {
			lastValue = bar
		}
	}
	for _, bar := range line {
		if unicode.IsDigit((bar)) {
			firstValue = bar
			break
		}
	}
	return string(firstValue) + string(lastValue)
}
