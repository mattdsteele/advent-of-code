package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	tst.Equals(t, 2, 2)
}

func TestParseInstruction(t *testing.T) {
	i := parseInstruction("1002")
	tst.Equals(t, 2, i.opcode)
	tst.Equals(t, []int{0, 1}, i.paramModes)
}

func TestParseOriginalInstruction(t *testing.T) {
	i := parseInstruction("3")
	tst.Equals(t, 3, i.opcode)
	tst.Equals(t, 0, len(i.paramModes))
}

func TestRunSimple(t *testing.T) {
	p := make("3,0,4,0,99", "1")
	tst.Equals(t, 0, p.position)
}

func TestTick(t *testing.T) {
	p := make("3,0,104,5,99", "1")
	p.tick()
	tst.Equals(t, 2, p.position)
	tst.Equals(t, "1", p.commands[0])

	p.tick()
	tst.Equals(t, 4, p.position)
	tst.Equals(t, []string{"5"}, p.output)

	p.tick()
	tst.Equals(t, 4, p.position)
}

func TestSample(t *testing.T) {
	p := make("1002,4,3,4,33", "1")
	p.tick()
	tst.Equals(t, "99", p.commands[4])
	tst.Equals(t, []string{"1002", "4", "3", "4", "99"}, p.commands)
}
func TestAnother(t *testing.T) {
	p := make("1001,4,3,4,33", "1")
	p.tick()
	tst.Equals(t, "36", p.commands[4])
	tst.Equals(t, 4, p.position)
}

func TestNegative(t *testing.T) {
	p := make("1101,100,-1,4,0", "1")
	p.tick()
	tst.Equals(t, "99", p.commands[4])
	tst.Equals(t, 4, p.position)
}
