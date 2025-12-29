package main

import (
	"math"
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var sample = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

var lines = util.SliceAtLine(sample)

func TestSilver(t *testing.T) {
	tst.Equals(t, math.Sqrt(48), distance(&Point{0, 0, 0}, &Point{4, 4, 4}))
	g := parse(lines)
	tst.Equals(t, 190, len(g.comparisons))

	// first value
	tst.Equals(t, g.comparisons[0].p1.x, 162)
	tst.Equals(t, g.comparisons[0].p2.x, 425)

	tst.Equals(t, g.comparisons[1].p1.x, 431)
	tst.Equals(t, g.comparisons[1].p2.x, 162)

	tst.Equals(t, 40, g.silver(10))
}

func TestGold(t *testing.T) {
	g := parse(lines)
	tst.Equals(t, 25272, g.gold())
}
