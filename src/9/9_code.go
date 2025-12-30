package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/9/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(lines []string) int {
	return parse(lines).silver()
}

func gold() {
	lines := util.ReadFile("src/9/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

type Game struct {
	points [][]int
}

func parse(lines []string) *Game {
	g := Game{}
	for _, l := range lines {
		strs := strings.Split(l, ",")
		xStr, yStr := strs[0], strs[1]
		xInt, _ := strconv.Atoi(xStr)
		yInt, _ := strconv.Atoi(yStr)
		g.points = append(g.points, []int{xInt, yInt})
	}
	return &g
}
func (g *Game) silver() int {
	max := 0
	for _, i := range g.points {
		for _, j := range g.points {
			a := area(i, j)
			if a > max {
				max = a
			}
		}
	}
	return max
}
func area(p1, p2 []int) int {
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]
	return (abs(x2, x1) + 1) * (abs(y2, y1) + 1)
}
func abs(x, y int) int {
	diff := y - x
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}
