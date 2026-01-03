package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sampleData = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestParse(t *testing.T) {
	lines := util.SliceAtLine(sampleData)
	p := parse(lines)
	tst.Equals(t, 50, p.silver())
}

func TestGold(t *testing.T) {
	lines := util.SliceAtLine(sampleData)
	p := parse(lines)
	tst.Equals(t, false, p.isInside(11, 1, 4, 4))
	tst.Equals(t, true, p.isInside(11, 1, 7, 3))
	tst.Equals(t, false, p.isInside(11, 1, 2, 3))

	tst.Equals(t, true, p.isInside(2, 5, 7, 3))
	tst.Equals(t, false, p.isInside(2, 5, 7, 1))

	tst.Equals(t, false, p.isInside(2, 5, 9, 7))
	tst.Equals(t, true, p.isInside(9, 5, 11, 7))

	tst.Equals(t, true, p.pointIsInside2(11, 7))
	tst.Equals(t, 24, p.gold())
}

func TestIntersects(t *testing.T) {
	// both same direction
	assertIntersects(t, 0, 2, 2, 2, 0, 1, 1, 1, false)
	assertIntersects(t, 2, 0, 2, 2, 1, 1, 1, 3, false)

	// overlaps
	assertIntersects(t, 2, 0, 2, 2, 1, 1, 3, 1, true)
	assertIntersects(t, 2, 0, 2, 2, 3, 1, 1, 1, true)
	assertIntersects(t, 2, 2, 2, 0, 3, 1, 1, 1, true)
	assertIntersects(t, 2, 2, 2, 0, 1, 1, 3, 1, true)
	assertIntersects(t, 3, 1, 1, 1, 2, 2, 2, 0, true)
	assertIntersects(t, 3, 1, 1, 1, 2, 0, 2, 2, true)

	// no overlap
	assertIntersects(t, 0, 0, 2, 0, 1, 3, 1, 1, false)
	assertIntersects(t, 2, 0, 0, 0, 1, 3, 1, 1, false)
}

func TestInside(t *testing.T) {
	lines := util.SliceAtLine(sampleData)
	p := parse(lines)
	tst.Equals(t, true, p.pointIsInside2(3, 4))
	tst.Equals(t, false, p.pointIsInside2(0, 0))
	tst.Equals(t, true, p.pointIsInside2(4, 3))
}

func assertIntersects(t *testing.T, p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y int, intersects bool) {
	segment1 := &Segment{&Point{p1x, p1y}, &Point{p2x, p2y}}
	segment2 := &Segment{&Point{p3x, p3y}, &Point{p4x, p4y}}
	tst.Equals(t, intersects, segmentsIntersect(segment1, segment2))
}
