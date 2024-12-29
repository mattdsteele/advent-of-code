package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

var results map[string]int = make(map[string]int)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/11/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	iterations := 25
	count := 0

	in := input[0]
	for _, i := range strings.Split(in, " ") {
		n, _ := strconv.Atoi(i)
		count += iterate(1, n, iterations)
	}
	return strconv.Itoa(count)
}

func gold() {
	lines := util.ReadFile("src/11/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	iterations := 75
	count := 0

	in := input[0]
	for _, i := range strings.Split(in, " ") {
		n, _ := strconv.Atoi(i)
		count += iterate(1, n, iterations)
	}
	return strconv.Itoa(count)
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

func iterate(count, n, iterationsRemaining int) int {
	key := fmt.Sprintf("%d,%d", n, iterationsRemaining)
	val := results[key]
	if val > 0 {
		return val
	}

	if iterationsRemaining == 0 {
		results[key] = count
		return count
	}

	iterationsRemaining--

	if n == 0 {
		c := iterate(count, 1, iterationsRemaining)
		results[key] = c
		return c
	}

	ns := strconv.Itoa(n)
	if len(ns)%2 == 0 {
		firstS := ns[0 : len(ns)/2]
		secondS := ns[len(ns)/2:]
		first, _ := strconv.Atoi(firstS)
		second, _ := strconv.Atoi(secondS)
		c := iterate(count, first, iterationsRemaining) + iterate(1, second, iterationsRemaining)
		results[key] = c
		return c
	}

	c := iterate(count, n*2024, iterationsRemaining)
	results[key] = c
	return c
}
