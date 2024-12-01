package main

import (
	"fmt"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	return "input"
}

func gold() {
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}
