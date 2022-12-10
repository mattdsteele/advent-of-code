package main

import (
	"fmt"
	"sort"
	"strconv"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/8/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) int {
	return parseBoard(input).visibleTrees()
}

func gold() {
	lines := util.ReadFile("src/8/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

func parseBoard(lines []string) *Board {
	b := new(Board)
	b.height = len(lines)
	b.width = len(lines[0])

	for x, l := range lines {
		for y, c := range l {
			height, _ := strconv.Atoi(string(c))
			t := new(Tree)
			t.height = height
			t.x = x
			t.y = y
			b.trees = append(b.trees, t)
		}
	}
	return b
}

type Board struct {
	trees  []*Tree
	width  int
	height int
}

func (b *Board) get(x, y int) *Tree {
	for _, t := range b.trees {
		if t.x == x && t.y == y {
			return t
		}
	}
	panic("could not find tree")
}
func (b *Board) visibleTrees() (count int) {
	x, y := 0, 0
	for x < b.width {
		for y < b.height {
			if b.visible(x, y) {
				count++
			}
			y++
		}
		x++
		y = 0
	}
	return count
}

func (b *Board) visible(x, y int) bool {
	if x == 0 || y == 0 {
		return true
	}
	if x+1 == b.width || y+1 == b.height {
		return true
	}

	t := b.get(x, y)
	return b.greatestHeightAround(x, y) < t.height
}

func (b *Board) greatestHeightAround(x, y int) int {
	heights := []int{b.greatestHeightAbove(x, y), b.greatestHeightBelow(x, y), b.greatestHeightLeft(x, y), b.greatestHeightRight(x, y)}
	sort.IntSlice.Sort(heights)
	return heights[0]
}
func (b *Board) greatestHeightAbove(x, y int) int {
	yToCheck := 0
	greatestHeight := 0
	for yToCheck < y {
		h := b.get(x, yToCheck).height
		if h > greatestHeight {
			greatestHeight = h
		}
		yToCheck++
	}
	return greatestHeight
}

func (b *Board) greatestHeightLeft(x, y int) int {
	xToCheck := 0
	greatestHeight := 0
	for xToCheck < x {
		h := b.get(xToCheck, y).height
		if h > greatestHeight {
			greatestHeight = h
		}
		xToCheck++
	}
	return greatestHeight
}

func (b *Board) greatestHeightRight(x, y int) int {
	xToCheck := x + 1
	greatestHeight := 0
	for xToCheck < b.width {
		h := b.get(xToCheck, y).height
		if h > greatestHeight {
			greatestHeight = h
		}
		xToCheck++
	}
	return greatestHeight
}

func (b *Board) greatestHeightBelow(x, y int) int {
	yToCheck := y + 1
	greatestHeight := 0
	for yToCheck < b.height {
		h := b.get(x, yToCheck).height
		if h > greatestHeight {
			greatestHeight = h
		}
		yToCheck++
	}
	return greatestHeight
}

type Tree struct {
	x      int
	y      int
	height int
}
