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
	fmt.Println(silverResult(util.ReadFile("./src/3/input.txt")))
}

func silverResult(lines []string) int {
	return gammaRate(lines) * epsilonRate(lines)
}

func epsilonRate(lines []string) int {
	var in string
	numChars := len(lines[0])
	for i := 0; i < numChars; i++ {
		zeroBits := getBits(lines, i, "0")
		oneBits := getBits(lines, i, "1")
		if zeroBits > oneBits {
			in += "0"
		} else {
			in += "1"
		}
	}
	intVal, _ := strconv.ParseInt(in, 2, 64)
	return int(intVal)
}

func co2Step(lines []string, i int) (nextLines []string) {
	strat := func(zeroCount, oneCount int) string {
		if oneCount >= zeroCount {
			return "0"
		} else {
			return "1"
		}
	}
	return getRemainingStrings(lines, i, strat)
}

type elToKeep func(zeroCount, oneCount int) string

func goldResult(lines []string) int {
	return co2Rating(lines) * oxygenRating(lines)
}

type step func(lines []string, index int) []string

func rating(lines []string, step step) int {
	var remainingLines []string
	for i := 0; ; i++ {
		remainingLines = step(lines, i)
		if len(remainingLines) == 1 {
			intVal, _ := strconv.ParseInt(remainingLines[0], 2, 64)
			return int(intVal)
		}
		lines = remainingLines
	}
}
func co2Rating(lines []string) int {
	return rating(lines, co2Step)
}

func oxygenRating(lines []string) int {
	return rating(lines, oxygenStep)
}

func oxygenStep(lines []string, idx int) []string {
	strat := func(zeroCount, oneCount int) string {
		if zeroCount > oneCount {
			return "0"
		} else {
			return "1"
		}
	}
	return getRemainingStrings(lines, idx, strat)
}

func getRemainingStrings(lines []string, idx int, elToKeepStrategy elToKeep) (remainingStrings []string) {
	zeroCount, oneCount := 0, 0
	var elToKeep string
	for _, line := range lines {
		el := line[idx : idx+1]
		if el == "0" {
			zeroCount++
		} else {
			oneCount++
		}
	}
	elToKeep = elToKeepStrategy(zeroCount, oneCount)
	for _, line := range lines {
		if elToKeep == line[idx:idx+1] {
			remainingStrings = append(remainingStrings, line)
		}
	}

	return remainingStrings
}

func getBits(lines []string, i int, s string) (total int) {
	for _, line := range lines {
		if line[i:i+1] == s {
			total++
		}
	}
	return total
}

func gammaRate(lines []string) int {
	var in string
	numChars := len(lines[0])
	for i := 0; i < numChars; i++ {
		zeroBits := getBits(lines, i, "0")
		oneBits := getBits(lines, i, "1")
		if zeroBits > oneBits {
			in += "1"
		} else {
			in += "0"
		}
	}
	intVal, _ := strconv.ParseInt(in, 2, 64)
	return int(intVal)
}

func gold() {
	fmt.Println(goldResult(util.ReadFile("./src/3/input.txt")))
}
