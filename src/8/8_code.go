package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/8/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	g := parse(input)
	return fmt.Sprintf("%d", g.silver(1000))
}

func gold() {
	lines := util.ReadFile("src/8/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	g := parse(input)
	return fmt.Sprintf("%d", g.gold())
}

type Point struct {
	x, y, z int
}

type Junctions struct {
	points map[*Point]bool
}

type Game struct {
	points      []*Point
	comparisons []*Comparison
	junctions   []*Junctions
}

func (g *Game) gold() interface{} {
	for _, c := range g.comparisons {
		g.tick(c)

		if len(g.points) == len(g.junctions[0].points) {
			// found the last two
			return c.p1.x * c.p2.x
		}
	}
	panic("could not find")
}

type Comparison struct {
	p1, p2   *Point
	distance float64
}

func (g *Game) silver(connections int) int {
	for i := 0; i < connections; i++ {
		comp := g.comparisons[i]
		g.tick(comp)
	}

	// now show the 3 largest circuits
	sort.Slice(g.junctions, func(p1, p2 int) bool {
		return len(g.junctions[p1].points) > len(g.junctions[p2].points)
	})
	items := 3
	total := 1
	for i := 0; i < items; i++ {
		size := len(g.junctions[i].points)
		total *= size
	}
	return total
}

func (g *Game) tick(comp *Comparison) {
	p1 := comp.p1
	p2 := comp.p2

	// check if p1 is in a set
	// add to set

	p1Junction := g.getJunction(p1)
	p2Junction := g.getJunction(p2)
	if p1Junction != nil && p2Junction != nil {

		if p2Junction != p1Junction {
			// merge sets
			for k := range p2Junction.points {
				p1Junction.points[k] = true
			}

			// delete original junction
			for i, j := range g.junctions {
				if j == p2Junction {
					g.junctions = append(g.junctions[:i], g.junctions[i+1:]...)
				}
			}
		} else {
		}
	} else if p1Junction == nil && p2Junction != nil {
		p2Junction.points[p1] = true

	} else if p2Junction == nil && p1Junction != nil {
		p1Junction.points[p2] = true
	} else {
		j := Junctions{}
		j.points = make(map[*Point]bool)
		j.points[p1] = true
		j.points[p2] = true
		g.junctions = append(g.junctions, &j)
	}
}

func (g *Game) getJunction(p1 *Point) *Junctions {
	for _, j := range g.junctions {
		if j.points[p1] == true {
			return j
		}
	}
	return nil
}

func parse(lines []string) *Game {
	g := Game{}
	g.points = []*Point{}
	for _, l := range lines {
		g.points = append(g.points, parseLine(l))
	}

	for _, p1 := range g.points {
		for _, p2 := range g.points {
			g.comparisons = append(g.comparisons, &Comparison{p1, p2, distance(p1, p2)})
		}
	}

	sort.Slice(g.comparisons, func(p1, p2 int) bool {
		return g.comparisons[p1].distance < g.comparisons[p2].distance
	})

	// remove dupes
	filteredComparisons := []*Comparison{}
	for i, c := range g.comparisons {
		if i%2 == 0 && c.distance > 0 {
			filteredComparisons = append(filteredComparisons, c)
		}
	}
	g.comparisons = filteredComparisons

	return &g
}

func parseLine(line string) *Point {
	pts := strings.Split(line, ",")
	xStr, yStr, zStr := pts[0], pts[1], pts[2]
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	z, _ := strconv.Atoi(zStr)
	return &Point{x, y, z}
}

func distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2) + math.Pow(float64(p2.z-p1.z), 2))
}
