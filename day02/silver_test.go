package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

type parse struct {
	Input         string
	ContainsTwo   bool
	ContainsThree bool
}

func TestParse(t *testing.T) {
	t.Log("Parse")
	tests := []parse{
		{"abcdef", false, false},
		{"bababc", true, true},
		{"abbcde", true, false},
		{"abcccd", false, true},
		{"aabcdd", true, false},
	}

	for _, test := range tests {
		containsTwo, containsThree := checkEntry(test.Input)
		if containsTwo != test.ContainsTwo {
			t.Fatal("Does not eq containsTwo", containsTwo, test.ContainsTwo, test.Input)
		}
		if containsThree != test.ContainsThree {
			t.Fatal("Does not eq containsThree", containsThree, test.ContainsThree, test.Input)
		}
	}
}
func TestChecksum(t *testing.T) {
	t.Log("Checksum")
	input := `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	checksum := getChecksum(util.SliceAtLine(input))
	expected := 12
	if checksum != expected {
		t.Fatal("Bad checksum", checksum, expected)
	}
}

