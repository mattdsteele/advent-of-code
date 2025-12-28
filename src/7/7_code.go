package main

import (
	"fmt"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/7/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	result := parse(input).silver()
	return fmt.Sprintf("%d", result)
}

func gold() {
	lines := util.ReadFile("src/7/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	return "input"
}

type Game struct {
	lines     []string
	tachLines []string
	beams     map[int]bool
}

func parse(lines []string) *Game {
	g := Game{}
	g.lines = lines
	g.tachLines = []string{}
	g.beams = make(map[int]bool)

	for i, l := range lines {
		if i%2 == 0 {
			g.tachLines = append(g.tachLines, l)
		}
	}
	return &g
}

func (g *Game) silverTick(line string) int {
	splits := 0
	for i, c := range line {
		if string(c) == "^" && g.beams[i] {
			g.beams[i] = false
			g.beams[i-1] = true
			g.beams[i+1] = true
			splits++
		}
		if string(c) == "S" {
			g.beams[i] = true
		}
	}
	return splits
}

func (g *Game) silver() int {
	splits := 0
	for _, l := range g.tachLines {
		splits += g.silverTick(l)
	}
	return splits
}
