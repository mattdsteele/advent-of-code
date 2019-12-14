package main

import (
	"fmt"
	"math"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type asteroid struct {
	x int
	y int
}

func (a asteroid) bearing(b asteroid) float64 {
	x := b.x - a.x
	y := b.y - a.y
	deg := math.Atan2(float64(x), float64(y)) / math.Pi * 180
	if deg < 0 {
		return 360 + deg
	}
	return deg
}

func (f field) visible(a asteroid) int {
	bearings := make(map[float64]bool)
	for _, c := range f.asteroids {
		if a != c {
			b := a.bearing(c)
			bearings[b] = true
		}
	}
	return len(bearings)
}

func (f field) mostVisible() (largest int, mostVisible asteroid) {
	for _, a := range f.asteroids {
		visible := f.visible(a)
		if visible > largest {
			largest = visible
			mostVisible = a
		}
	}
	return largest, mostVisible
}

type field struct {
	asteroids []asteroid
}

func main() {
	silver()
}

func silver() {
	f := util.ReadFile("./src/10/input.txt")
	i, _ := parseLines(f).mostVisible()
	fmt.Println(i)
}

func parseLines(lines []string) *field {
	f := new(field)
	for y, l := range lines {
		for x, a := range strings.Split(l, "") {
			if a == "#" {
				f.asteroids = append(f.asteroids, asteroid{x, y})
			}
		}
	}
	return f
}

func parse(in string) *field {
	lines := util.SliceAtLine(in)
	return parseLines(lines)
}
