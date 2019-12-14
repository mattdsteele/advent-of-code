package main

import (
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
	tst.Equals(t, float64(0), d.bearing(c))
	tst.Equals(t, float64(45), d.bearing(b))
	tst.Equals(t, float64(90), d.bearing(a))
	tst.Equals(t, float64(270), d.bearing(e))
	tst.Equals(t, float64(180), d.bearing(f))
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
