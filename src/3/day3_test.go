package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestMove(t *testing.T) {
	w := newWire("")
	w.move("R234")
	tst.Equals(t, path{234, 0}, w.position)
	tst.Equals(t, len(w.paths), 234)
}

func TestMove2(t *testing.T) {
	w := newWire("")
	w.move("U12")
	w.move("L22")
	tst.Equals(t, path{-22, 12}, w.position)
}

func TestParse(t *testing.T) {
	parsed := parse("R234,U12,L22")
	tst.Equals(t, []string{"R234", "U12", "L22"}, parsed)
}

func TestWire(t *testing.T) {
	def := newWire("R8,U5,L5,D3")
	tst.Equals(t, path{3, 2}, def.position)
}

func TestIntersections(t *testing.T) {
	wire1 := newWire("R8,U5,L5,D3")
	wire2 := newWire("U7,R6,D4,L4")
	i := intersections(wire1, wire2)
	tst.Equals(t, 2, len(i))
}
func TestManhattan(t *testing.T) {
	tst.Equals(t, 6, manhattanDistance(path{3, 3}))
	tst.Equals(t, 7, manhattanDistance(path{4, -3}))
}

func TestClosestDistance(t *testing.T) {
	wire1 := newWire("R8,U5,L5,D3")
	wire2 := newWire("U7,R6,D4,L4")
	tst.Equals(t, 6, closestDistance(wire1, wire2))
	tst.Equals(t, 159, closestDistance(newWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"), newWire("U62,R66,U55,R34,D71,R55,D58,R83")))
	tst.Equals(t, 135, closestDistance(newWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"), newWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")))
}

func TestStepsTo(t *testing.T) {
	wire1 := newWire("R8,U5,L5,D3")
	tst.Equals(t, 20, wire1.stepsTo(path{3, 3}))
	wire2 := newWire("U7,R6,D4,L4")
	tst.Equals(t, 20, wire2.stepsTo(path{3, 3}))
}

func TestBestSteps(t *testing.T) {
	tst.Equals(t, 610, fewestCombinedSteps(newWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"), newWire("U62,R66,U55,R34,D71,R55,D58,R83")))
	tst.Equals(t, 410, fewestCombinedSteps(newWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"), newWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")))

}
