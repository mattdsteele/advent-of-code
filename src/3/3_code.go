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
	largest := 0
	size := 12
	for _, c := range candidates(input, size) {
		candidate, _ := strconv.Atoi(c)
		if candidate > largest {
			largest = candidate
		}
	}
	return largest
}

func candidates(input string, size int) []string {
	candidateMap := map[string]bool{}
	cans := &candidateMap
	builtCandidate := ""
	for i, s := range input {
		remaining := input[i+1:]
		builtCandidate = ""
		recurCandidates(builtCandidate+string(s), remaining, size, cans)

	}

	candidateList := []string{}
	for k, _ := range candidateMap {
		candidateList = append(candidateList, k)
	}

	return candidateList
}

func recurCandidates(builtCandidate string, remainingInput string, size int, candidates *map[string]bool) {
	for i, s := range remainingInput {
		remaining := remainingInput[i+1:]
		if len(builtCandidate)+1 == size {
			candMap := *candidates
			candMap[builtCandidate+string(s)] = true
		} else {
			recurCandidates(builtCandidate+string(s), remaining, size, candidates)
		}
		// append
	}

}
