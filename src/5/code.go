package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/5/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	return parse(input).playSilver()
}

func gold() {
	lines := util.ReadFile("src/5/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

type Game struct {
	columns   []*Column
	idxOfBase int
	steps     []*Step
}

type Step struct {
	count int
	from  int
	to    int
}

func (g *Game) playSilver() string {
	for _, s := range g.steps {
		c := s.count
		for c > 0 {
			v := g.columns[s.from-1].pop()
			g.columns[s.to-1].push(v)
			c--
		}
	}
	return g.printBase()
}
func (g *Game) printBase() (b string) {
	for _, c := range g.columns {
		b += c.items[0]
	}
	return b
}

type Column struct {
	items []string
}

func (c *Column) pop() string {
	item := c.items[0]
	c.items = c.items[1:]
	return item
}

func (c *Column) push(s string) {
	c.items = append([]string{s}, c.items...)
}

func parse(lines []string) *Game {
	g := new(Game)
	g.columns = []*Column{}
	idxOfBase := indexOfBase(lines)
	g.idxOfBase = idxOfBase
	base := lines[idxOfBase]
	countOfColumns := 0
	for _, i := range base {
		if i != ' ' {
			countOfColumns++
		}
	}

	for countOfColumns > 0 {
		g.columns = append(g.columns, new(Column))
		countOfColumns--
	}

	for i := 0; i < idxOfBase; i++ {
		for j, r := range lines[i] {
			if !strings.ContainsRune("[] ", r) {
				idxOfColumn := (j - 1) / 4
				column := g.columns[idxOfColumn]
				column.items = append(column.items, string(r))
			}
		}
	}

	for _, s := range lines[idxOfBase+2:] {
		step := parseStep(s)
		g.steps = append(g.steps, step)
	}

	return g
}
func indexOfBase(lines []string) int {
	for i, l := range lines {
		if strings.IndexRune(l, '1') == 1 {
			return i
		}
	}
	panic("could not find index")
}

func parseStep(l string) *Step {
	s := new(Step)
	r, _ := regexp.Compile("move (\\d*) from (\\d*) to (\\d*)")
	matches := r.FindStringSubmatch(l)
	count, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])
	s.count = count
	s.from = from
	s.to = to
	return s
}
