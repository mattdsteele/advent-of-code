package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	val1, val2 := parse("128392-643281")
	tst.Equals(t, "128392", val1)
	tst.Equals(t, "643281", val2)
}
func TestDoubleNumber(t *testing.T) {
	fixture := new(adjacentEqual)
	tst.Equals(t, true, fixture.passes("111111"))
	tst.Equals(t, false, fixture.passes("123456"))
	tst.Equals(t, true, fixture.passes("123466"))
	tst.Equals(t, true, fixture.passes("223450"))
}

func TestDigitsDontDecrease(t *testing.T) {
	fixture := new(digitsDontDecrease)
	tst.Equals(t, true, fixture.passes("111111"))
	tst.Equals(t, true, fixture.passes("123789"))
	tst.Equals(t, false, fixture.passes("223450"))
}
func TestAllRulesMet(t *testing.T) {
	fixture := initSilverRules()
	tst.Equals(t, true, fixture.passes("111111"))
	tst.Equals(t, false, fixture.passes("123789"))
	tst.Equals(t, false, fixture.passes("223450"))
}

func TestRange(t *testing.T) {
	silverInput := "128392-643281"
	numPassingInputs := passingSilverPasswords(silverInput)
	tst.Equals(t, 2050, numPassingInputs)
}

func TestExactlyTwoRepeatingDigits(t *testing.T) {
	fixture := new(exactlyTwoRepeatingDigits)
	tst.Equals(t, true, fixture.passes("112233"))
	tst.Equals(t, false, fixture.passes("111234"))
	tst.Equals(t, false, fixture.passes("011134"))
	tst.Equals(t, false, fixture.passes("011114"))
	tst.Equals(t, true, fixture.passes("001114"))
	tst.Equals(t, false, fixture.passes("123444"))
	tst.Equals(t, true, fixture.passes("111122"))
}
