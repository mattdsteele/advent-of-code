package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	test "github.com/mattdsteele/advent-of-code/testing"
)

type parse struct {
	Input  string
	Parsed []string
}

func TestFuel(t *testing.T) {
	t.Log("Fuel")
	type input struct {
		Input    string
		Expected int
	}
	tests := []input{
		{"12", 2},
		{"14", 2},
		{"1969", 654},
		{"100756", 33583},
	}
	for _, test := range tests {
		fuel := fuel(test.Input)
		if test.Expected != fuel {
			t.Fatal("Sums not eq", test.Expected, fuel)
		}
	}
}
func TestRecursiveFuel(t *testing.T) {
	t.Log("Fuel Recursive")
	type input struct {
		Input    int
		Expected int
	}
	tests := []input{
		{12, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, test := range tests {
		fuel := recurseFuel(test.Input)
		if test.Expected != fuel {
			t.Fatal("Sums not eq", test.Expected, fuel)
		}
	}
}

func TestFuelSum(t *testing.T) {
	t.Log("FuelSum")
	f := fuelSum(util.SliceAtLine(`12
14
1969
100756`))
	test.Equals(t, 2+2+654+33583, f)
}
