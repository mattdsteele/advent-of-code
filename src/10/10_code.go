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
	lines := util.ReadFile("src/10/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) int {
	return game(input).silver()
}

func gold() {
	lines := util.ReadFile("src/10/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) int {
	return 0
}

type Machine struct {
	buttons        []bool
	goal           []bool
	wiringDiagrams []*WiringDiagram
}

func (m *Machine) silver() int {
	max := 0
	nextStepStates := make(map[string]bool)
	origStates := [][]bool{m.buttons}
	for {
		max++
		for _, origState := range origStates {
			match, nss := buttonTick(origState, m)
			if match {
				return max
			}
			for _, ns := range nss {
				nextStepStates[str(ns)] = true
			}
		}

		// nothing found, move to next round
		origStates = [][]bool{}
		for k := range nextStepStates {
			origStates = append(origStates, decomp(k))
		}
	}
}

func game(lines []string) *Game {
	g := Game{}
	for _, l := range lines {
		g.machines = append(g.machines, parse(l))
	}
	return &g
}

func decomp(state string) (marshalled []bool) {
	for _, s := range state {
		if s == '#' {
			marshalled = append(marshalled, true)
		} else {
			marshalled = append(marshalled, false)
		}
	}
	return marshalled
}

func buttonTick(state []bool, m *Machine) (match bool, newStates [][]bool) {
	for _, wd := range m.wiringDiagrams {
		ns := push(state, wd.buttons)
		if matches(ns, m.goal) {
			return true, newStates
		}
		newStates = append(newStates, ns)
	}
	return false, newStates
}

func matches(state, goal []bool) bool {
	for i := range state {
		if state[i] != goal[i] {
			return false
		}
	}
	return true
}

func push(state []bool, button []int) []bool {
	newState := []bool{}
	for _, s := range state {
		newState = append(newState, s)
	}
	for _, b := range button {
		newState[b] = !state[b]
	}
	return newState
}

func str(state []bool) string {
	st := ""
	for _, s := range state {
		if s {
			st += "#"
		} else {
			st += "."
		}
	}
	return st
}

type WiringDiagram struct {
	buttons []int
}

type Game struct {
	machines []*Machine
}

func (g *Game) silver() int {
	total := 0
	for _, m := range g.machines {
		total += m.silver()
	}
	return total
}

func parse(line string) *Machine {
	machine := Machine{}
	buttons := regexp.MustCompile(`\[(.*)\]`)
	b := buttons.FindStringSubmatch(line)[1]
	for _, g := range b {
		machine.buttons = append(machine.buttons, false)
		machine.goal = append(machine.goal, g == '#')
	}

	w := regexp.MustCompile(`\(([\d,]*)\)`).FindAllStringSubmatch(line, -1)
	for _, b := range w {
		machine.wiringDiagrams = append(machine.wiringDiagrams, parseWiringDiagram(b[1]))
	}

	return &machine
}

func parseWiringDiagram(buttons string) *WiringDiagram {
	wd := WiringDiagram{}
	b := strings.Split(buttons, ",")
	for _, i := range b {
		j, _ := strconv.Atoi(i)
		wd.buttons = append(wd.buttons, j)
	}
	return &wd
}
