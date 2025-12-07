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
	lines := util.ReadFile("src/3/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	sum := 0
	for _, i := range input {
		sum += silverJoltage(i)
	}
	return strconv.Itoa(sum)
}

func gold() {
	lines := util.ReadFile("src/3/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	var sum int64 = 0
	for _, i := range lines {
		sum += goldJoltage(i)
	}
	return strconv.FormatInt(sum, 10)
}

func silverJoltage(input string) int {
	return int(candidates(input, 2))
}

func goldJoltage(input string) int64 {
	size := 12
	return candidates(input, size)
}

func candidates(input string, size int) int64 {
	var largest int64 = 0
	for i, s := range input {
		remaining := input[i+1:]
		recurCandidates(string(s), remaining, size, &largest)

	}

	return largest
}

func recurCandidates(builtCandidate string, remainingInput string, size int, largest *int64) {
	for i, s := range remainingInput {
		remaining := remainingInput[i+1:]
		newCand := builtCandidate + string(s)
		if len(builtCandidate)+1 == size {
			candidate, _ := strconv.ParseInt(newCand, 10, 64)
			if candidate > *largest {
				*largest = candidate
			}
		} else {
			// if there is a chance it can win, recur it
			maxVal := newCand + strings.Repeat("9", size-len(newCand))
			largestOpt, _ := strconv.ParseInt(maxVal, 10, 64)
			if largestOpt < *largest {
				continue
			}

			recurCandidates(newCand, remaining, size, largest)
		}
	}

}
