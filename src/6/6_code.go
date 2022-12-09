package main

import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/6/input.txt")
	fmt.Println(silverCalculate(lines[0]))
}

func silverCalculate(input string) int {
	return uniqueForValue(input, 4)
}

func uniqueForValue(input string, count int) int {
	signal := ""
	for i, r := range input {
		signal += string(r)
		if len(signal) > count {
			signal = signal[1:]
			if allUnique(signal) {
				return i + 1
			}
		}
	}
	panic("failed to find unique vals")
}

func allUnique(signal string) bool {
	for _, c := range signal {
		if strings.Count(signal, string(c)) > 1 {
			return false
		}
	}
	return true
}

func gold() {
	lines := util.ReadFile("src/6/input.txt")
	fmt.Println(goldCalculate(lines[0]))
}

func goldCalculate(lines string) int {
	return uniqueForValue(lines, 14)
}
