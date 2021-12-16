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
	in := util.ReadFile("src/5/input.txt")
	fmt.Println(setup(in).overlapping())
}

func gold() {
	// Print gold result
}

type line struct {
	from *point
	to   *point
}

type point struct {
	x int
	y int
}

type grid struct {
	lines        []*line
	pointMapping map[point]int
}

func setup(input []string) *grid {
	g := &grid{}
	for _, in := range input {
		g.lines = append(g.lines, parseLine(in))
	}
	return g
}

func (g *grid) setupGrid() {
	g.pointMapping = make(map[point]int)
	for _, line := range g.lines {
		points := toRange(line)
		for _, point := range points {
			g.append(point)
		}
	}
}

func (g *grid) append(point *point) {
	g.pointMapping[*point] = g.pointMapping[*point] + 1
}

func (g *grid) greaterThan(input int) int {
	totalGreater := 0
	for _, v := range g.pointMapping {
		if v > input {
			totalGreater++
		}
	}
	return totalGreater
}

func (g *grid) overlapping() int {
	g.setupGrid()
	return g.greaterThan(1)
}

func parseLine(def string) *line {
	points := strings.Split(def, " -> ")
	from := parsePoint(points[0])
	to := parsePoint(points[1])
	return &line{from, to}
}

func parsePoint(def string) *point {
	defs := strings.Split(def, ",")
	x, _ := strconv.Atoi(defs[0])
	y, _ := strconv.Atoi(defs[1])
	return &point{x, y}
}

func toRange(line *line) (r []*point) {
	if line.from.x == line.to.x {
		delta := findRange(line.from.y, line.to.y)
		for _, d := range delta {
			r = append(r, &point{line.from.x, d})
		}
	} else if line.from.y == line.to.y {
		delta := findRange(line.from.x, line.to.x)
		for _, d := range delta {
			r = append(r, &point{d, line.from.y})
		}
	}
	return r
}

func findRange(from, to int) (r []int) {
	var min, max int
	if from < to {
		min = from
		max = to
	} else {
		min = to
		max = from
	}
	for i := min; i <= max; i++ {
		r = append(r, i)
	}
	return r
}
