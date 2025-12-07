package main

import (
	"fmt"
	"strconv"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/4/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	return strconv.Itoa(parse(input).silver())
}

func gold() {
	lines := util.ReadFile("src/4/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

type Point struct {
	row    int
	column int
}

type Game struct {
	points map[string]*Point
}

func key(row, column int) string {
	return fmt.Sprintf("%d,%d", row, column)
}

func (g Game) pointAt(row, column int) *Point {
	return g.points[key(row, column)]

}
func (g Game) surrounding(row, column int) (total int) {
	for r := row - 1; r <= row+1; r++ {
		for c := column - 1; c <= column+1; c++ {
			if !(r == row && c == column) {
				candidate := g.pointAt(r, c)
				if candidate != nil {
					total++
				}
			}
		}
	}
	return total
}
func (g Game) surroundingPoints(point *Point) (total int) {
	return g.surrounding(point.row, point.column)
}
func (g Game) silver() (total int) {
	threshold := 4
	for _, v := range g.points {
		if g.surroundingPoints(v) < threshold {
			total++
		}
	}
	return total
}

func parse(input []string) *Game {
	g := Game{}
	g.points = make(map[string]*Point)

	for r, row := range input {
		for c, item := range row {
			if item == '@' {
				p := Point{}
				p.row = r
				p.column = c
				g.points[key(r, c)] = &p
			}
		}
	}
	return &g
}
