package main

import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("./src/8/input.txt")
	g := makeGame(lines)
	fmt.Println(g.uniqueSegments())
}

func gold() {
	// Print gold result
}

type Game struct {
	lines []*Line
}

func makeGame(inputs []string) *Game {
	g := Game{}
	for _, i := range inputs {
		g.lines = append(g.lines, parseLine(i))
	}
	return &g
}

func (g *Game) uniqueSegments() int {
	segments := 0
	uniqueSegmentLengths := []int{2, 4, 3, 7}
	for _, l := range g.lines {
		for _, i := range l.output {
			if contains(uniqueSegmentLengths, len(i)) {
				segments++
			}
		}
	}
	return segments
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Line struct {
	raw            string
	signalPatterns []string
	output         []string
}

func parseLine(input string) *Line {
	l := Line{}
	l.raw = input
	split := strings.Split(input, "|")
	l.signalPatterns = signalPatternsFromRaw(split[0])
	l.output = outputFromRaw(split[1])
	return &l
}

func signalPatternsFromRaw(s string) []string {
	return []string{}
}

func outputFromRaw(s string) []string {
	trimmed := strings.TrimSpace(s)
	values := strings.Split(trimmed, " ")
	return values
}
