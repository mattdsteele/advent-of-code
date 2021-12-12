package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	silver()
	gold()
}

func silver() {
	fmt.Println("Silver result")

}

func gold() {
	// Print gold result
}

type board struct {
	rows []*row
}
type row struct {
	cols []*entry
}
type entry struct {
	value    int
	selected bool
}

type game struct {
	draws  []int
	boards []*board
}

func parse(lines []string) *game {
	g := new(game)
	steps := lines[0]
	for _, step := range strings.Split(steps, ",") {
		num, _ := strconv.Atoi(step)
		g.draws = append(g.draws, num)
	}

	lineSize := len(lines)
	for lineSize > ((len(g.boards) * 6) + 2) {
		board := parseBoard(lines, len(g.boards))
		g.boards = append(g.boards, board)
	}
	return g
}

func parseBoard(lines []string, boardNumber int) (b *board) {
	b = new(board)
	b.rows = make([]*row, 0)
	for i := 0; i < 5; i++ {
		idx := boardNumber*6 + 2 + i
		splitLines := strings.Split(lines[idx], " ")
		r := new(row)
		nums := make([]*entry, 0)
		for _, l := range splitLines {
			trimmed := strings.TrimSpace(l)
			if trimmed != "" {
				count, _ := strconv.Atoi(trimmed)
				nums = append(nums, &entry{count, false})
			}
		}
		r.cols = nums
		b.rows = append(b.rows, r)
	}
	return b
}
