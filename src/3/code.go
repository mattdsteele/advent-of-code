package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
	"github.com/zoumo/goset"
)

type path struct {
	x int
	y int
}

type wire struct {
	paths    []path
	position path
}

func newWire(guide string) *wire {
	w := new(wire)
	w.position = path{0, 0}

	commands := parse(guide)
	for _, command := range commands {
		w.move(command)
	}
	return w
}

func (w *wire) move(command string) {
	dir := command[0]
	res := command[1:]
	amount, _ := strconv.Atoi(res)
	w.doMove(dir, amount)
}
func (w *wire) doMove(dir byte, numLeft int) {
	if numLeft == 0 {
		return
	}
	if dir == 'R' {
		w.position.x++
	}
	if dir == 'L' {
		w.position.x--
	}
	if dir == 'U' {
		w.position.y++
	}
	if dir == 'D' {
		w.position.y--
	}
	w.paths = append(w.paths, path{w.position.x, w.position.y})
	w.doMove(dir, numLeft-1)
}
func parse(input string) []string {
	if input == "" {
		return []string{}
	}
	return strings.Split(input, ",")
}

func intersections(w1, w2 *wire) (is []path) {
	w1Paths := goset.NewSetFrom(w1.paths)
	w2Paths := goset.NewSetFrom(w2.paths)
	crossings := w1Paths.Intersect(w2Paths)
	for _, i := range goset.SetToSlice.Elements(crossings) {
		is = append(is, i.(path))
	}
	return is
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func manhattanDistance(p path) int {
	return abs(p.x) + abs(p.y)
}
func closestDistance(w1, w2 *wire) (closest int) {
	is := intersections(w1, w2)
	for _, i := range is {
		d := manhattanDistance(i)
		if closest == 0 || d < closest {
			closest = d
		}
	}
	return closest
}

func main() {
	lines := util.ReadFile("src/3/input.txt")
	w1 := newWire(lines[0])
	w2 := newWire(lines[1])
	fmt.Println(closestDistance(w1, w2))
}
