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
