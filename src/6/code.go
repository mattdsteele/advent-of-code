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
	fmt.Println(g.fishCount())
}

func gold() {
	line := util.ReadFile("./src/6/input.txt")[0]
	g := parse(line)
	g.playRounds(256)
	fmt.Println(g.fishCount())
}

type Game struct {
	rounds           int
	fishCountPerItem map[int]int
}

func parse(initialState string) *Game {
	strStates := strings.Split(initialState, ",")
	g := new(Game)
	g.fishCountPerItem = make(map[int]int)

	for _, s := range strStates {
		i, _ := strconv.Atoi(s)
		g.fishCountPerItem[i] = g.fishCountPerItem[i] + 1
	}
	return g
}

func (g *Game) tick() {
	g.rounds++
	num0 := g.fishCountPerItem[0]
	num1 := g.fishCountPerItem[1]
	num2 := g.fishCountPerItem[2]
	num3 := g.fishCountPerItem[3]
	num4 := g.fishCountPerItem[4]
	num5 := g.fishCountPerItem[5]
	num6 := g.fishCountPerItem[6]
	num7 := g.fishCountPerItem[7]
	num8 := g.fishCountPerItem[8]

	g.fishCountPerItem[7] = num8
	g.fishCountPerItem[6] = num7 + num0
	g.fishCountPerItem[5] = num6
	g.fishCountPerItem[4] = num5
	g.fishCountPerItem[3] = num4
	g.fishCountPerItem[2] = num3
	g.fishCountPerItem[1] = num2
	g.fishCountPerItem[0] = num1
	g.fishCountPerItem[8] = num0
}

func (g *Game) playRounds(times int) {
	for i := 0; i < times; i++ {
		g.tick()
	}
}

func (g *Game) fishCount() int {
	totalCount := 0
	for _, v := range g.fishCountPerItem {
		totalCount += v
	}
	return totalCount
}
