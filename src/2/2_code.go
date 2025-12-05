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
	line := input[0]
	ranges := strings.Split(line, ",")
	invalidCount := 0

	for _, idRange := range ranges {
		els := strings.Split(idRange, "-")
		first, _ := strconv.Atoi(els[0])
		last, _ := strconv.Atoi(els[1])
		for i := first; i <= last; i++ {
			if !valid(i) {
				invalidCount += i
			}

		}
	}

	return strconv.Itoa(invalidCount)
}

func valid(input int) bool {
	inputStr := strconv.Itoa(input)
	strLen := len(inputStr)
	if strLen%2 != 0 {
		return true
	}
	first := inputStr[0 : strLen/2]
	second := inputStr[(strLen / 2):]
	return first != second
}

func goldValid(input int) bool {
	inputStr := strconv.Itoa(input)
	for inputSize := 1; inputSize <= len(inputStr)/2; inputSize++ {
		if len(inputStr)%inputSize != 0 {
			continue
		}

		first := inputStr[0:inputSize]
		allEqual := true
		for i := 1; i < len(inputStr)/inputSize; i++ {
			next := inputStr[i*inputSize : i*inputSize+inputSize]
			if first != next {
				allEqual = false
				break
			}
		}
		if allEqual {
			return false
		}
	}
	return true
}

func gold() {
	lines := util.ReadFile("src/2/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	line := input[0]
	ranges := strings.Split(line, ",")
	invalidCount := 0

	for _, idRange := range ranges {
		els := strings.Split(idRange, "-")
		first, _ := strconv.Atoi(els[0])
		last, _ := strconv.Atoi(els[1])
		for i := first; i <= last; i++ {
			if !goldValid(i) {
				invalidCount += i
			}

		}
	}

	return strconv.Itoa(invalidCount)
}
