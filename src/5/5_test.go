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

func TestEq8(t *testing.T) {
	p := make("3,9,8,9,10,9,4,9,99,-1,8", "8")
	p.runUntilDone()
	tst.Equals(t, []string{"1"}, p.output)
	p = make("3,9,8,9,10,9,4,9,99,-1,8", "9")
	p.runUntilDone()
	tst.Equals(t, []string{"0"}, p.output)
}

func TestLt8(t *testing.T) {
	p := make("3,9,7,9,10,9,4,9,99,-1,8", "7")
	p.runUntilDone()
	tst.Equals(t, []string{"1"}, p.output)
	p = make("3,9,7,9,10,9,4,9,99,-1,8", "9")
	p.runUntilDone()
	tst.Equals(t, []string{"0"}, p.output)
}
func TestImmediateEq8(t *testing.T) {
	p := make("3,3,1108,-1,8,3,4,3,99", "8")
	p.runUntilDone()
	tst.Equals(t, "1", p.output[0])
	p = make("3,3,1108,-1,8,3,4,3,99", "9")
	p.runUntilDone()
	tst.Equals(t, "0", p.output[0])
	p = make("3,3,1108,-1,8,3,4,3,99", "7")
	p.runUntilDone()
	tst.Equals(t, "0", p.output[0])
}
func TestImmediateLt8(t *testing.T) {
	p := make("3,3,1107,-1,8,3,4,3,99", "5")
	p.runUntilDone()
	tst.Equals(t, "1", p.output[0])
	p = make("3,3,1107,-1,8,3,4,3,99", "8")
	p.runUntilDone()
	tst.Equals(t, "0", p.output[0])
	p = make("3,3,1107,-1,8,3,4,3,99", "9")
	p.runUntilDone()
	tst.Equals(t, "0", p.output[0])
}

func TestJmp(t *testing.T) {
	p := make("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "0")
	p.runUntilDone()
	tst.Equals(t, "0", p.output[0])
	p = make("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "1")
	p.runUntilDone()
	tst.Equals(t, "1", p.output[0])
	p = make("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "-1")
	p.runUntilDone()
	tst.Equals(t, "1", p.output[0])
}

func TestLargerEx(t *testing.T) {
	in := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	p := make(in, "7")
	p.runUntilDone()
	tst.Equals(t, "999", p.output[0])

}
