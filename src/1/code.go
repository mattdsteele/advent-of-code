package main

import (
	"fmt"
	"strconv"

	util "github.com/mattdsteele/advent-of-code"
)

func fuel(input string) int {
	in, _ := strconv.Atoi(input)
	return fuelCost(in)
}
func fuelCost(input int) int {
	return (input / 3) - 2
}

func recurseFuel(input int) int {
	fuelCost := fuelCost(input)
	if fuelCost <= 0 {
		return 0
	}
	return fuelCost + recurseFuel(fuelCost)
}

func fuelSum(ins []string) int {
	total := 0
	for _, in := range ins {
		total = total + fuel(in)
	}
	return total
}

func main() {
	silver()
	gold()
}

func silver() {
	sum := fuelSum(util.ReadFile("./src/1/input.txt"))
	fmt.Println(sum)
}
func gold() {
	ins := util.ReadFile("./src/1/input.txt")
	total := 0
	for _, in := range ins {
		f, _ := strconv.Atoi(in)
		total += recurseFuel(f)
	}
	fmt.Println(total)
}
