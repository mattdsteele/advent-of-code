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
	lines := util.ReadFile("src/11/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	iterations := 25
	in := input[0]
	for i := 0; i < iterations; i++ {
		in = silverIterate(in)
	}
	return strconv.Itoa(len(strings.Split(in, " ")))
}

func gold() {
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func silverIterate(exampleInput string) string {
	vals := []string{}
	for _, i := range strings.Split(exampleInput, " ") {
		j, _ := strconv.Atoi(i)
		if j == 0 {
			vals = append(vals, "1")
			continue
		}

		if len(i)%2 == 0 {
			firstS := i[0 : len(i)/2]
			secondS := i[len(i)/2:]
			first, _ := strconv.Atoi(firstS)
			second, _ := strconv.Atoi(secondS)
			vals = append(vals, strconv.Itoa(first))
			vals = append(vals, strconv.Itoa(second))
			continue
		}

		vals = append(vals, strconv.Itoa(j*2024))

	}
	return strings.Join(vals, " ")
}
