package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	t.Log("Parse")
	parsed := parse("1,0,0,0,99")
	util.Assert(t, util.SlicesEqual([]int{1, 0, 0, 0, 99}, parsed), "values not equal")
}

func TestOpcode(t *testing.T) {
	t.Log("Silver opcode")
	util.Assert(t, util.SlicesEqual([]int{2, 0, 0, 0, 99}, opcode(parse("1,0,0,0,99"))), "values not equal")
	util.Assert(t, util.SlicesEqual([]int{2, 3, 0, 6, 99}, opcode(parse("2,3,0,3,99"))), "values not equal")
	util.Assert(t, util.SlicesEqual([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, opcode(parse("1,1,1,4,99,5,6,0,99"))), "values not equal")
}

func TestFixAlarm(t *testing.T) {
	util.Assert(t, util.SlicesEqual([]int{1, 12, 2, 0, 99}, fixAlarm(parse("1,0,0,0,99"), 12, 2)), "values not equal")
}
