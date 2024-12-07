package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/2/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	safeCount := 0
	for _, l := range input {
		if silverSafe(l) {
			safeCount++
		}
	}
	return strconv.Itoa(safeCount)
}

func gold() {
	lines := util.ReadFile("src/2/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	safeCount := 0
	for _, l := range input {
		if goldSafe(l) {
			safeCount++
		}
	}
	return strconv.Itoa(safeCount)
}

func goldSafe(line string) bool {
	numi := toIntSlice(line)
	// Original value?
	firstCheck := safeChecks(numi)
	if firstCheck {
		return true
	}

	// If not, remove each entry and perform a check
	for i := range numi {
		freshNum := toIntSlice(line)
		if safeChecks(removeEntry(freshNum, i)) {
			return true
		}
	}
	return false
}

func removeEntry(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func silverSafe(line string) bool {
	numi := toIntSlice(line)
	return safeChecks(numi)
}

func toIntSlice(line string) []int {
	nums := strings.Split(line, " ")
	numi := []int{}
	for _, num := range nums {
		i, _ := strconv.Atoi(num)
		numi = append(numi, i)
	}
	return numi
}

func safeChecks(numi []int) bool {
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
