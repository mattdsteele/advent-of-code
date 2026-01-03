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

func goldCalculate(lines []string) int {
	return parse(lines).gold()
}

type Game struct {
	points   [][]int
	segments []*Segment
}

// https://wrfranklin.org/Research/Short_Notes/pnpoly.html
func (g *Game) pointIsInside(x, y int) bool {
	isInside := false
	for _, s := range g.segments {
		i := s.point1
		j := s.point2
		if (i.y > y) != (j.y > y) &&
			x < (j.x-i.x)*(y-i.y)/(j.y-i.y)+i.x {
			isInside = !isInside
		}
	}
	return isInside
}

func between(p, a, b int) bool {
	return p >= a && p <= b || p <= a && p >= b
}

func (g *Game) pointIsInside2(x, y int) bool {
	inside := false
	P := &Point{x, y}
	for _, s := range g.segments {
		A := s.point1
		B := s.point2

		if P.x == A.x && P.y == A.y || P.x == B.x && P.y == B.y {
			return true
		}
		if A.y == B.y && P.y == A.y && between(P.x, A.x, B.x) {
			return true
		}

		if between(P.y, A.y, B.y) { // if P inside the vertical range
			// filter out "ray pass vertex" problem by treating the line a little lower
			if P.y == A.y && B.y >= A.y || P.y == B.y && A.y >= B.y {
				continue
			}
			// calc cross product `PA X PB`, P lays on left side of AB if c > 0
			c := (A.x-P.x)*(B.y-P.y) - (B.x-P.x)*(A.y-P.y)
			if c == 0 {
				return true
			}
			if (A.y < B.y) == (c > 0) {
				inside = !inside
			}
		}
	}
	return inside
}

/*
// Source - https://stackoverflow.com/a
// Posted by timepp, modified by community. See post 'Timeline' for change history
// Retrieved 2026-01-02, License - CC BY-SA 4.0

* Get relationship between a point and a polygon using ray-casting algorithm
 * @param {{x:number, y:number}} P: point to check
 * @param {{x:number, y:number}[]} polygon: the polygon
 * @returns -1: outside, 0: on edge, 1: inside
function relationPP(P, polygon) {
    const between = (p, a, b) => p >= a && p <= b || p <= a && p >= b
    let inside = false
    for (let i = polygon.length-1, j = 0; j < polygon.length; i = j, j++) {
        const A = polygon[i]
        const B = polygon[j]
        // corner cases
        if (P.x == A.x && P.y == A.y || P.x == B.x && P.y == B.y) return 0
        if (A.y == B.y && P.y == A.y && between(P.x, A.x, B.x)) return 0

        if (between(P.y, A.y, B.y)) { // if P inside the vertical range
            // filter out "ray pass vertex" problem by treating the line a little lower
            if (P.y == A.y && B.y >= A.y || P.y == B.y && A.y >= B.y) continue
            // calc cross product `PA X PB`, P lays on left side of AB if c > 0
            const c = (A.x - P.x) * (B.y - P.y) - (B.x - P.x) * (A.y - P.y)
            if (c == 0) return 0
            if ((A.y < B.y) == (c > 0)) inside = !inside
        }
    }

    return inside? 1 : -1
}

*/

func (g *Game) isInside(ix, iy, jx, jy int) bool {
	top := Segment{&Point{ix, iy}, &Point{ix, jy}}
	right := Segment{&Point{ix, jy}, &Point{jx, jy}}
	bottom := Segment{&Point{jx, jy}, &Point{jx, iy}}
	left := Segment{&Point{jx, iy}, &Point{ix, iy}}
	box := []*Segment{&top, &right, &bottom, &left}
	for _, bs := range box {
		for _, gs := range g.segments {
			if segmentsIntersect(bs, gs) {
				return false
			}
		}
	}

	// check each corner
	for _, bs := range box {
		p := bs.point1
		if !g.pointIsInside2(p.x, p.y) {
			return false
		}
	}

	// check one inside
	minX := min(ix, jx)
	minY := min(iy, jy)
	if !g.pointIsInside2(minX+1, minY+1) {
		return false
	}

	return true
}

func (g *Game) gold() int {
	max := 0
	for _, i := range g.points {
		for _, j := range g.points {
			a := area(i, j)
			if a > 0 {
				if g.isInside(i[0], i[1], j[0], j[1]) && a > max {
					// fmt.Printf("new max (%d): %d,%d - %d,%d\n", a, i[0], i[1], j[0], j[1])
					max = a
				}
			}
		}
	}
	return max
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

	// create segments
	var prevPoint *Point
	for _, p := range g.points {
		currPoint := &Point{p[0], p[1]}
		if prevPoint != nil {
			g.segments = append(g.segments, &Segment{prevPoint, currPoint})
		}
		prevPoint = currPoint
	}
	// wrap around
	first := g.points[0]
	g.segments = append(g.segments, &Segment{prevPoint, &Point{first[0], first[1]}})
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

type Point struct {
	x, y int
}

type Segment struct {
	point1 *Point
	point2 *Point
}

func (s *Segment) minX() int {
	return min(s.point1.x, s.point2.x)
}
func (s *Segment) maxX() int {
	return max(s.point1.x, s.point2.x)
}
func (s *Segment) minY() int {
	return min(s.point1.y, s.point2.y)
}
func (s *Segment) maxY() int {
	return max(s.point1.y, s.point2.y)
}

func (s *Segment) isVertical() bool {
	return s.point1.x == s.point2.x
}

func segmentsIntersect(s1 *Segment, s2 *Segment) bool {
	if s1.isVertical() == s2.isVertical() {
		// both are same direction, they do not intersect
		return false
	}
	// 2,0 - 2,2
	// 1,1 - 3,1
	// if there is overlap, they intersect

	// x overlap scenario 1
	if s1.maxX() > s2.minX() && s1.minX() < s2.maxX() {
		if s1.maxY() > s2.minY() && s1.minY() < s2.maxY() {
			// y overlap scenario 1
			return true
		}

		if s2.maxY() > s1.minY() && s2.minY() < s1.maxY() {
			// y overlap scenario 2
			return true
		}
	}

	// x overlap scenario 1
	if s2.maxX() > s1.minX() && s2.minX() < s1.maxX() {
		if s1.maxY() > s2.minY() && s1.minY() < s2.maxY() {
			// y overlap scenario 1
			return true
		}

		if s2.maxY() > s1.minY() && s2.minY() < s1.maxY() {
			// y overlap scenario 2
			return true
		}

	}
	return false
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
