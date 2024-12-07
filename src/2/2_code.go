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
	lines := util.ReadFile("src/2/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	safeCount := 0
	for _, l := range input {
		if isSafe(l) {
			safeCount++
		}
	}
	return strconv.Itoa(safeCount)
}

func gold() {
	lines := util.ReadFile("src/2/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func isSafe(line string) bool {
	nums := strings.Split(line, " ")
	numi := []int{}
	for _, num := range nums {
		i, _ := strconv.Atoi(num)
		numi = append(numi, i)
	}
	if alwaysDecreasing(numi) {
		return true
	}
	if alwaysIncreasing(numi) {
		return true
	}
	return false
}

func alwaysIncreasing(numi []int) bool {
	last := 0
	for i, n := range numi {
		if i == 0 {
			last = n
		} else if last >= n {
			return false
		} else if n-last > 3 {
			return false
		}
		last = n
	}
	return true
}
func alwaysDecreasing(numi []int) bool {
	last := 0
	for i, n := range numi {
		if i == 0 {
			last = n
		} else if last <= n {
			return false
		} else if last-n > 3 {
			return false
		}
		last = n
	}
	return true
}
