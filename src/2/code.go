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
	fmt.Println(goldResult(util.ReadFile("src/2/input.txt")))
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

func goldResult(lines []string) (total int) {
	for _, line := range lines {
		total += goldPoint(line)
	}
	return total
}
func goldPoint(line string) int {
	pointsMapping := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	vals := strings.Split(line, " ")
	me := vals[1]
	totalPoints := pointsMapping[me]

	foe := vals[0]
	foeTranslate := equivalentsMap[foe]
	// which one requires it
	if totalPoints == 3 {
		totalPoints += pointsMap[foeTranslate]
	} else {
		if totalPoints == 0 {
			// which throw is defeated by foeTranslate?
			toThrow := defeatsMap[foeTranslate]
			totalPoints += pointsMap[toThrow]
		} else {
			// which throw defeats foeTranslate?
			options := []string{"X", "Y", "Z"}
			for _, option := range options {
				if defeatsMap[option] == foeTranslate {
					totalPoints += pointsMap[option]
				}
			}
		}
	}
	return totalPoints
}
