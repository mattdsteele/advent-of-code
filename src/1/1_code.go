package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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
	left := []int{}
	right := []int{}
	for _, line := range input {
		l, r := getEntries(line)
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)
	totalDistance := 0
	for i, l := range left {
		r := right[i]
		totalDistance += distance(l, r)
	}
	return strconv.Itoa(totalDistance)
}

func gold() {
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func distance(from, to int) int {
	if from > to {
		return from - to
	}
	return to - from
}

func getEntries(input string) (left, right int) {
	values := strings.Split(input, "   ")
	left, _ = strconv.Atoi(values[0])
	right, _ = strconv.Atoi(values[1])
	return left, right
}
