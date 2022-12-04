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
	lines := util.ReadFile("src/4/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) (count int) {
	for _, i := range input {
		if enclosed(i) {
			count++
		}
	}
	return count
}

func gold() {
	lines := util.ReadFile("src/4/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

type assignment struct {
	low  int
	high int
}

func makeAssignment(s string) *assignment {
	a := new(assignment)

	values := strings.Split(s, "-")
	low, _ := strconv.Atoi(values[0])
	high, _ := strconv.Atoi(values[1])
	a.high = high
	a.low = low

	return a
}
func enclosed(s string) bool {
	assignments := strings.Split(s, ",")
	assignment1 := makeAssignment(assignments[0])
	assignment2 := makeAssignment(assignments[1])
	assignment1Lower := assignment1.low >= assignment2.low && assignment1.high <= assignment2.high
	assignment2Lower := assignment2.low >= assignment1.low && assignment2.high <= assignment1.high
	return assignment1Lower || assignment2Lower
}
