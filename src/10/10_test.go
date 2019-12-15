package main

import (
	"math"
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	in := `.#..#
.....
#####
....#
...##`
	field := parse(in)
	tst.Equals(t, 10, len(field.asteroids))
	tst.Equals(t, field.asteroids[0], asteroid{1, 0})
}

func TestDegrees(t *testing.T) {
	a := asteroid{1, 0}
	b := asteroid{1, 1}
	c := asteroid{0, 1}
	d := asteroid{0, 0}
	e := asteroid{-1, 0}
	f := asteroid{0, -1}
	tst.Equals(t, float64(180), d.bearing(c))
	tst.Equals(t, float64(135), d.bearing(b))
	tst.Equals(t, float64(90), d.bearing(a))
	tst.Equals(t, float64(270), d.bearing(e))
	tst.Equals(t, float64(0), d.bearing(f))
}

func TestAsteroidsVisible(t *testing.T) {
	in := `.#..#
.....
#####
....#
...##`
	field := parse(in)
	first := field.asteroids[0]
	tst.Equals(t, 7, field.visible(first))
}

func TestMostDetected(t *testing.T) {
	in := `.#..#
.....
#####
....#
...##`
	field := parse(in)
	num, a := field.mostVisible()
	tst.Equals(t, 8, num)
	tst.Equals(t, field.asteroids[8], a)
}

func TestMoreExamples(t *testing.T) {
	tests := []struct {
		in            string
		x             int
		y             int
		asteroidsSeen int
	}{
		{`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`, 5, 8, 33},
		{`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`, 1, 2, 35},
	}
	for _, fixture := range tests {
		f := parse(fixture.in)
		num, a := f.mostVisible()
		tst.Equals(t, num, fixture.asteroidsSeen)
		tst.Equals(t, a.x, fixture.x)
		tst.Equals(t, a.y, fixture.y)
	}
}

func astAt(f *field, x, y int) asteroid {
	for _, a := range f.asteroids {
		if a.x == x && a.y == y {
			return a
		}
	}
	panic("failed")
}

func TestDistance(t *testing.T) {
	a := asteroid{0, 0}
	b := asteroid{1, 1}
	c := asteroid{0, 1}
	tst.Equals(t, float64(1), a.distance(c))
	tst.Equals(t, float64(1), c.distance(a))
	tst.Equals(t, math.Sqrt(2), a.distance(b))
}
func TestDataStructure(t *testing.T) {
	fld := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	f := parse(fld)
	a := astAt(f, 8, 3)
	entries := f.bearingMappings(a)
	tst.Assert(t, entries != nil, "should exist")
	zeros := entries[float64(0)]
	tst.Equals(t, 2, len(zeros))
}

func TestBearingsDistance(t *testing.T) {
	fld := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	f := parse(fld)
	a := astAt(f, 8, 3)
	entries := f.bearingMappings(a)
	sorted := sortMappings(entries, a)
	tst.Equals(t, astAt(f, 8, 1), sorted[0][0])
	tst.Equals(t, astAt(f, 9, 0), sorted[1][0])
}

func TestXthVaporized(t *testing.T) {
	fld := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	f := parse(fld)
	a := astAt(f, 11, 13)
	tst.Equals(t, astAt(f, 11, 12), f.vaporized(a, 1))
	tst.Equals(t, astAt(f, 10, 16), f.vaporized(a, 100))
	tst.Equals(t, astAt(f, 11, 1), f.vaporized(a, 299))
}
