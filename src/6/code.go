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

func (s *system) orbitCount() (count int) {
	for _, o := range s.objects {
		count += o.orbitCount()
	}
	return count
}

func main() {
	silver()
}

func silver() {
	sys := new(system)
	input := util.ReadFile("./src/6/input.txt")
	for _, line := range input {
		sys.parseOrbit(line)
	}
	fmt.Println(sys.orbitCount())
}
