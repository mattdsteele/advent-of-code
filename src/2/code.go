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
	fmt.Println(silverData(lines))
}

func silverData(sampleData []string) int {
	var horizontal, vertical int
	for _, step := range sampleData {
		direction, amt := parseLine(step)
		switch direction {
		case "forward":
			horizontal += amt
		case "up":
			vertical -= amt
		case "down":
			vertical += amt
		}
	}
	return horizontal * vertical
}

func parseLine(line string) (direction string, value int) {
	values := strings.Split(line, " ")
	direction = values[0]
	value, _ = strconv.Atoi(values[1])
	return direction, value
}

func gold() {
	// Print gold result
}
