package main
import (
	"reflect"
	"testing"
)

type parse struct {
	Input  string
	Parsed []string
}

func TestParse(t *testing.T) {
	t.Log("Parse")
	tests := []parse{
		{
			`+1
+1
+1`, []string{"+1", "+1", "+1"}},
		{
			`-1
+15
+100`, []string{"-1", "+15", "+100"}},
	}

	for _, test := range tests {
		parsed := toSlice(test.Input)
		if !slicesEqual(parsed, test.Parsed) {
			t.Fatal("Slices not eq", parsed, test.Parsed)
		}

	}
}

func TestSum(t *testing.T) {
	t.Log("Sum")
	type input struct {
		Input    []string
		Expected int
	}
	tests := []input{
		{[]string{"+1", "+1", "+1"}, 3},
	}
	for _, test := range tests {
		sum := add(test.Input)
		if test.Expected != sum {
			t.Fatal("Sums not eq", test.Expected, sum)
		}
	}
}
func slicesEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

