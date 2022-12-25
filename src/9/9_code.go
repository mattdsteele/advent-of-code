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

func silverCalculate(input []string) string {
	g := newGame(input)
	g.runLines()
	return fmt.Sprintf("%d", g.audit.numberOfSpots())
}

func gold() {
	lines := util.ReadFile("src/9/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

type Knot struct {
	x, y int
}

func (k *Knot) serialize() string {
	return fmt.Sprintf("%d,%d", k.x, k.y)
}

type Game struct {
	instructions []string
	head, tail   *Knot
	audit        *Audit
}

func (g *Game) logPosition(k *Knot) {
	g.audit.locations[k.serialize()] = true
}
func (g *Game) moveHead(direction string) {
	switch direction {
	case "L":
		g.head.x = g.head.x - 1
	case "R":
		g.head.x = g.head.x + 1
	case "U":
		g.head.y = g.head.y + 1
	case "D":
		g.head.y = g.head.y - 1
	}
}
func (g *Game) runLine(line string) {
	vals := strings.Split(line, " ")
	direction := vals[0]
	count, _ := strconv.Atoi(vals[1])
	for count > 0 {
		g.playTurn(direction)
		count--
	}
}
func (g *Game) runLines() {
	for _, i := range g.instructions {
		g.runLine(i)
	}
}

func (g *Game) moveTail() {
	h := g.head
	t := g.tail
	// in same row
	if h.x == t.x {
		if abs(h.y-t.y) > 1 {
			if h.y > t.y {
				t.y++
			} else {
				t.y--
			}
		}
		// in same column
	} else if h.y == t.y {
		if abs(h.x-t.x) > 1 {
			if h.x > t.x {
				t.x++
			} else {
				t.x--
			}
		}
		// diagonal but not touching
	} else if abs(h.x-t.x) > 1 || abs(h.y-t.y) > 1 {
		if h.y > t.y {
			t.y++
		} else {
			t.y--
		}
		if h.x > t.x {
			t.x++
		} else {
			t.x--
		}
	}
}
func (g *Game) playTurn(direction string) {
	g.moveHead(direction)
	g.moveTail()
	g.logPosition(g.tail)
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func newGame(instructions []string) *Game {
	g := new(Game)
	g.instructions = instructions
	g.head = new(Knot)
	g.tail = new(Knot)
	g.audit = newAudit()
	return g
}

type Audit struct {
	locations map[string]bool
}

func newAudit() *Audit {
	a := new(Audit)
	a.locations = make(map[string]bool)
	return a
}

func (a *Audit) numberOfSpots() int {
	return len(a.locations)
}
