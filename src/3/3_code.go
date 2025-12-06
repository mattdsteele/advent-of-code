package main

import (
	"fmt"
	"strconv"

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
	sum := 0
	for _, i := range lines {
		sum += goldJoltage(i)
	}
	return strconv.Itoa(sum)
}

func silverJoltage(input string) int {
	largest := 0
	for i, s := range input {
		substr := input[i+1:]
		for _, q := range substr {
			v := fmt.Sprintf("%s%s", string(s), string(q))
			candidate, _ := strconv.Atoi(v)
			if candidate > largest {
				largest = candidate
			}
		}
	}
	return largest
}

func goldJoltage(input string) int {
	size := 12
	return candidates(input, size)
}

func candidates(input string, size int) int {
	largest := 0
	builtCandidate := ""
	for i, s := range input {
		remaining := input[i+1:]
		builtCandidate = ""
		recurCandidates(builtCandidate+string(s), remaining, size, &largest)

	}

	return largest
}

func recurCandidates(builtCandidate string, remainingInput string, size int, largest *int) {
	for i, s := range remainingInput {
		remaining := remainingInput[i+1:]
		newCand := builtCandidate + string(s)
		candidate, _ := strconv.Atoi(newCand)
		if len(builtCandidate)+1 == size {
			if candidate > *largest {
				*largest = candidate
			}
		} else {
			recurCandidates(newCand, remaining, size, largest)
		}
	}

}
