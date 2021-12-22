package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	line := util.ReadFile("./src/7/input.txt")[0]
	fmt.Println(parse(line).bestFuelCost())
}

func gold() {
	// Print gold result
}

type Game struct {
	crabPositions []int
}

func (g *Game) min() (minValue int) {
	minValue = 999
	for _, i := range g.crabPositions {
		if i < minValue {
			minValue = i
		}
	}
	return minValue
}
func (g *Game) max() (maxValue int) {
	for _, i := range g.crabPositions {
		if i > maxValue {
			maxValue = i
		}
	}
	return maxValue
}

func (g *Game) fuelCost(level int) (totalFuel int) {
	for _, c := range g.crabPositions {
		totalFuel += fuel(level, c)
	}
	return totalFuel
}

func (g *Game) bestFuelCost() int {
	bestFuel := -1
	min := g.min()
	max := g.max()
	for i := min; i < max; i++ {
		fc := g.fuelCost(i)
		if bestFuel < 0 || fc < bestFuel {
			bestFuel = fc
		}
	}
	return bestFuel
}

func fuel(from, to int) int {
	total := from - to
	if total < 0 {
		return -total
	}
	return total
}

func parse(input string) *Game {
	g := Game{}
	strs := strings.Split(input, ",")
	for _, str := range strs {
		i, _ := strconv.Atoi(str)
		g.crabPositions = append(g.crabPositions, i)
	}
	return &g

}
