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
	line := util.ReadFile("./src/6/input.txt")[0]
	g := parse(line)
	g.playRounds(80)
	fmt.Print(g.fishCount())
}

func gold() {
	// Print gold result
}

type Game struct {
	rounds int
	fishes []int
}

func parse(initialState string) *Game {
	strStates := strings.Split(initialState, ",")
	g := new(Game)
	g.fishes = make([]int, 0)

	for _, s := range strStates {
		i, _ := strconv.Atoi(s)
		g.fishes = append(g.fishes, i)
	}
	return g
}

func (g *Game) tick() {
	g.rounds++
	numFishToAdd := 0
	for i, f := range g.fishes {
		f--
		if f == -1 {
			f = 6
			numFishToAdd++
		}
		g.fishes[i] = f
	}
	for i := 0; i < numFishToAdd; i++ {
		g.fishes = append(g.fishes, 8)
	}
}

func (g *Game) playRounds(times int) {
	for i := 0; i < times; i++ {
		g.tick()
	}
}

func (g *Game) fishCount() int {
	return len(g.fishes)
}
