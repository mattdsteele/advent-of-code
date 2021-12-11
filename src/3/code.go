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

func oxygenStep(lines []string, idx int) (remainingStrings []string) {
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
	if zeroCount > oneCount {
		elToKeep = "0"
	} else {
		elToKeep = "1"
	}
	for _, line := range lines {
		if elToKeep == line[idx:idx+1] {
			remainingStrings = append(remainingStrings, line)
		}
	}

	return remainingStrings
}

func oxygenRating(lines []string) int {
	var remainingLines []string
	for i := 0; ; i++ {
		remainingLines = oxygenStep(lines, i)
		if len(remainingLines) == 1 {
			intVal, _ := strconv.ParseInt(remainingLines[0], 2, 64)
			return int(intVal)
		}
		lines = remainingLines
	}
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
	// Print gold result
}
