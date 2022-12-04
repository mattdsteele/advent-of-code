package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestConvertRucksack(t *testing.T) {
	tst.Equals(t, 2, 2)
	exampleInput := `vJrwpWtwJgWrhcsFMMfFFhFp`
	sacks := getRucksacks(exampleInput)
	tst.Equals(t, 2, len(sacks))
	tst.Equals(t, "vJrwpWtwJgWr", sacks[0])
	tst.Equals(t, "hcsFMMfFFhFp", sacks[1])
	tst.Equals(t, "p", sharedItem(sacks))

	tst.Equals(t, "L", sharedItem(getRucksacks("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")))
}

func TestPriorityCount(t *testing.T) {
	tst.Equals(t, 1, priority("a"))
	tst.Equals(t, 2, priority("b"))
	tst.Equals(t, 27, priority("A"))
	tst.Equals(t, 42, priority("P"))
	tst.Equals(t, 38, priority("L"))
}

func TestSilver(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	lines := util.SliceAtLine(input)
	tst.Equals(t, 157, silverCalculate(lines))
}

func TestGoldSimilar(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`
	inputs := util.SliceAtLine(input)
	tst.Equals(t, "r", goldSimilar(inputs))
}

func TestGoldScore(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	lines := util.SliceAtLine(input)
	tst.Equals(t, 70, goldCalculate(lines))

}
