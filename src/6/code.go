package main

import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type object struct {
	orbiting *object
	name     string
}

func (o *object) orbitCount() int {
	if o.orbiting == nil {
		return 0
	}
	return 1 + o.orbiting.orbitCount()
}
func (o *object) isOrbiting(t *object) bool {
	if o == t {
		return true
	}
	if o.orbiting == nil {
		return false
	}
	return o.orbiting.isOrbiting(t)
}

type system struct {
	objects []*object
}

func (s *system) parseOrbit(line string) {
	lines := strings.Split(line, ")")
	orbitedName := lines[0]
	orbiterName := lines[1]
	orbited := s.find(orbitedName)
	orbiter := s.find(orbiterName)
	orbiter.orbiting = orbited
}

func (s *system) find(name string) *object {
	for _, o := range s.objects {
		if o.name == name {
			return o
		}
	}
	newObject := &object{nil, name}
	s.objects = append(s.objects, newObject)
	return newObject
}
func (s *system) youOrbiting() *object {
	return s.find("YOU").orbiting
}
func (s *system) sanOrbiting() *object {
	return s.find("SAN").orbiting
}
func (s *system) common(o1, o2 *object) *object {
	test := o1.orbiting
	for test != nil {
		if o2.isOrbiting(test) {
			return test
		}
		test = test.orbiting
	}
	return nil
}

func (s *system) orbitCount() (count int) {
	for _, o := range s.objects {
		count += o.orbitCount()
	}
	return count
}
func (s *system) hops(orbiter, orbiting *object) (count int) {
	test := orbiter
	for test != orbiting {
		count++
		test = test.orbiting
	}
	return count
}

func main() {
	silver()
	gold()
}

func makeSys() *system {
	sys := new(system)
	input := util.ReadFile("./src/6/input.txt")
	for _, line := range input {
		sys.parseOrbit(line)
	}
	return sys
}

func silver() {
	fmt.Println(makeSys().orbitCount())
}

func gold() {
	fmt.Println(goldTest(makeSys()))
}

func goldTest(s *system) int {
	common := s.common(s.youOrbiting(), s.sanOrbiting())
	count := s.hops(s.youOrbiting(), common) + s.hops(s.sanOrbiting(), common)
	return count
}
