package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var exampleInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var lines = util.SliceAtLine(exampleInput)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
	g := newGame(lines)
	tst.Equals(t, 8, len(g.instructions))
	tst.Equals(t, 0, g.head.x)
	tst.Equals(t, 0, g.head.y)
	tst.Equals(t, 0, g.tail.x)
	tst.Equals(t, 0, g.tail.y)
}

func TestAuditPositions(t *testing.T) {
	g := newGame(lines)
	k := new(Knot)
	k.x = 5
	k.y = 6
	g.logPosition(k)
	tst.Equals(t, 1, g.audit.numberOfSpots())

	k.x = 6
	g.logPosition(k)
	tst.Equals(t, 2, g.audit.numberOfSpots())

	k.x = 5
	g.logPosition(k)
	tst.Equals(t, 2, g.audit.numberOfSpots())
}

func TestMoveHead(t *testing.T) {
	g := newGame(lines)
	g.moveHead("L")
	tst.Equals(t, -1, g.head.x)
	tst.Equals(t, 0, g.head.y)

	g.moveHead("U")
	tst.Equals(t, -1, g.head.x)
	tst.Equals(t, 1, g.head.y)

	g.moveHead("D")
	tst.Equals(t, -1, g.head.x)
	tst.Equals(t, 0, g.head.y)

	g.moveHead("R")
	tst.Equals(t, 0, g.head.x)
	tst.Equals(t, 0, g.head.y)
}

func TestMoveTail(t *testing.T) {
	g := newGame(lines)
	g.moveTail()
	tst.Equals(t, "0,0", g.tail.serialize())

	g.moveHead("U")
	g.moveTail()
	tst.Equals(t, "0,0", g.tail.serialize())

	g.moveHead("U")
	g.moveTail()
	tst.Equals(t, "0,1", g.tail.serialize())

	g.head.x = 2
	g.head.y = 1
	g.tail.x = 0
	g.tail.y = 0
	g.moveTail()
	tst.Equals(t, "1,1", g.tail.serialize())
}

func TestRunLine(t *testing.T) {
	g := newGame(lines)
	g.runLine(lines[0])
	tst.Equals(t, "4,0", g.head.serialize())
	tst.Equals(t, "3,0", g.tail.serialize())
	tst.Equals(t, 4, g.audit.numberOfSpots())
}

func TestSilver(t *testing.T) {
	g := newGame(lines)
	g.runLines()
	tst.Equals(t, 13, g.audit.numberOfSpots())
}

func TestRunTurn(t *testing.T) {
	g := newGame(lines)
	g.playTurn("D")
	g.playTurn("D")
	tst.Equals(t, "0,-1", g.tail.serialize())
	tst.Equals(t, 2, g.audit.numberOfSpots())

	g.playTurn("U")
	g.playTurn("R")
	g.playTurn("R")
	tst.Equals(t, "1,-1", g.tail.serialize())
	tst.Equals(t, 3, g.audit.numberOfSpots())

	g.playTurn("L")
	g.playTurn("L")
	g.playTurn("L")
	tst.Equals(t, "0,-1", g.tail.serialize())
	tst.Equals(t, 3, g.audit.numberOfSpots())
}

func TestKnotFunctions(t *testing.T) {
	k := new(Knot)
	tst.Equals(t, "0,0", k.serialize())
	k.x = 12
	k.y = -4
	tst.Equals(t, "12,-4", k.serialize())
}
