package main
import (
	"strings"
	"testing"
)

func split(input string) []string {
	return strings.Split(input, ",")
}
func TestFreqCount(t *testing.T) {
	t.Log("Sum")
	type input struct {
		Input    []string
		Expected int
	}
	tests := []input{
		{[]string{"-1", "+1"}, 0},
		{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{split("-6,+3,+8,+5,-6"), 5},
		{split("+7,+7,-2,-7,-4"), 14},
	}
	for _, test := range tests {
		firstFreq := firstFrequency(test.Input)
		if test.Expected != firstFreq {
			t.Fatal("first freq not eq", test.Expected, firstFreq)
		}
	}
}

