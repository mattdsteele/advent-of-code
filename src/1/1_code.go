package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	// silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	left, right := parseAndSort(input)
	totalDistance := 0
	for i, l := range left {
		r := right[i]
		totalDistance += distance(l, r)
	}
	return strconv.Itoa(totalDistance)
}

func parseAndSort(input []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range input {
		l, r := getEntries(line)
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func gold() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	left, right := parseAndSort(lines)
	count := 0
	for _, l := range left {
		count += (l * numberEntries(right, l))
	}
	return strconv.Itoa(count)
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

func numberEntries(entries []int, check int) (count int) {
	for _, r := range entries {
		if r == check {
			count++
		}
	}
	return count
}
