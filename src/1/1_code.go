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
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	status := 50
	count := 0
	for _, i := range input {
		z := delta(i)
		status += z
		if status%100 == 0 {
			count++
		}
		if status >= 100 {
			status -= 100
		}
		if status <= 0 {
			status += 100
		}
	}
	return strconv.Itoa(count)
}

func delta(i string) int {
	val, _ := strconv.Atoi(i[1:])
	diff := i[0:1]
	if diff == "L" {
		val *= -1
	}
	return val
}

func gold() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}
