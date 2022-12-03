package main

import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

var equivalentsMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}
var defeatsMap = map[string]string{
	"X": "Z",
	"Y": "X",
	"Z": "Y",
}

var pointsMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	silver()
	gold()
}

func silver() {
	fmt.Println(silverResult(util.ReadFile("src/2/input.txt")))
}

func gold() {
	// Print gold result
}

func silverResult(lines []string) (total int) {
	for _, line := range lines {
		total += silverPoints(line)
	}
	return total
}

func silverPoints(line string) int {
	vals := strings.Split(line, " ")
	me := vals[1]

	throwPoints := pointsMap[me]

	foe := vals[0]
	foeTranslate := equivalentsMap[foe]

	//score points
	if me == foeTranslate {
		throwPoints += 3
	} else if defeatsMap[me] == foeTranslate {
		throwPoints += 6
	}

	return throwPoints
}
