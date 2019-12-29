package main

import (
	"fmt"
	util "github.com/mattdsteele/advent-of-code"
	"strconv"
	"strings"
)

type bond struct {
	amount   int
	chemical string
}
type pairing struct {
	target   *bond
	requires []*bond
}
type system struct {
	pairings []*pairing
}
type deps struct {
	bonds []*bond
}

func parseReqs(in string) *system {
	s := system{}
	s.pairings = make([]*pairing, 0)
	for _, st := range strings.Split(in, "\n") {
		s.pairings = append(s.pairings, parseLine(st))
	}
	return &s
}

func (s *system) oreCost() int {
	fuelVal := s.chemical("FUEL")
	distanceLevel := s.distanceFromOre("FUEL")
	bonds := []*bond{fuelVal.target}
	for distanceLevel > 0 {
		bonds = s.depsAtLevel(bonds, distanceLevel)
		distanceLevel--
	}
	if len(bonds) > 1 {
		panic(fmt.Sprintf("too many bonds found: %+v", bonds))
	}
	return bonds[0].amount
}

func (s *system) convert(b *bond) (converted []*bond) {
	ch := s.chemical(b.chemical)
	reqPerConversion := ch.target.amount
	numNeeded := b.amount
	conversionsRequired := numNeeded / reqPerConversion
	if numNeeded%reqPerConversion > 0 {
		conversionsRequired++
	}
	for _, r := range ch.requires {
		newR := &bond{r.amount * conversionsRequired, r.chemical}
		converted = append(converted, newR)
	}
	return converted
}
func (s *system) depsAtLevel(bonds []*bond, level int) (newDeps []*bond) {
	for _, b := range bonds {
		dist := s.distanceFromOre(b.chemical)
		if dist >= level {
			converted := s.convert(b)
			for _, c := range converted {
				newDeps = add(newDeps, c)
			}
		} else {
			newDeps = add(newDeps, b)
		}
	}
	return newDeps
}

func add(deps []*bond, b *bond) []*bond {
	for _, bond := range deps {
		if bond.chemical == b.chemical {
			bond.amount += b.amount
			return deps
		}
	}
	deps = append(deps, b)
	return deps
}

func (s *system) distanceFromOre(ch string) int {
	if ch == "ORE" {
		return 0
	}
	c := s.chemical(ch)
	maxDistance := 0
	for _, st := range c.requires {
		distance := s.distanceFromOre(st.chemical)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return 1 + maxDistance
}

func (s *system) chemical(ch string) *pairing {
	for _, p := range s.pairings {
		if p.target.chemical == ch {
			return p
		}
	}
	panic("could not find chemical " + ch)
}
func (s *system) chemicalsAtDistance(distance int) (chems []string) {
	for _, p := range s.pairings {
		if s.distanceFromOre(p.target.chemical) == distance {
			chems = append(chems, p.target.chemical)
		}
	}
	return chems
}

func toBond(in string) *bond {
	tar := strings.Split(in, " ")
	amt, _ := strconv.Atoi(tar[0])
	target := new(bond)
	target.amount = amt
	target.chemical = tar[1]
	return target
}
func parseLine(in string) *pairing {
	p := pairing{}
	st := strings.Split(in, " => ")
	p.target = toBond(st[1])
	p.requires = make([]*bond, 0)

	bonds := strings.Split(st[0], ", ")
	for _, b := range bonds {
		p.requires = append(p.requires, toBond(b))
	}
	return &p
}
func main() {
	silver()
}

func silver() {
	lines := util.ReadFile("./src/14/input.txt")
	s := parseReqs(strings.Join(lines, "\n"))
	fmt.Println(s.oreCost())
}
