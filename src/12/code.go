package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type moon struct {
	x  int
	y  int
	z  int
	dx int
	dy int
	dz int
}

type system struct {
	moons []*moon
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
func (m *moon) energy() int {
	pot := abs(m.x) + abs(m.y) + abs(m.z)
	kin := abs(m.dx) + abs(m.dy) + abs(m.dz)
	return pot * kin
}

func (s *system) energy() (e int) {
	for _, m := range s.moons {
		e += m.energy()
	}
	return e
}

func (s *system) add(m *moon) {
	s.moons = append(s.moons, m)
}

func (s *system) run(times int) {
	i := 0
	for i < times {
		s.tick()
		i++
	}
}

func (s *system) tick() {
	for _, m := range s.moons {
		for _, o := range s.moons {
			if m != o {
				m.gravity(o)
			}
		}
	}
	for _, m := range s.moons {
		m.velocity()
	}
}

func (m *moon) velocity() {
	m.x += m.dx
	m.y += m.dy
	m.z += m.dz
}

func (m *moon) gravity(o *moon) {
	if o.x > m.x {
		m.dx++
	} else if o.x < m.x {
		m.dx--
	}
	if o.y > m.y {
		m.dy++
	} else if o.y < m.y {
		m.dy--
	}
	if o.z > m.z {
		m.dz++
	} else if o.z < m.z {
		m.dz--
	}
}

func parse(line string) *moon {
	pos := strings.Split(line, ",")
	xs := strings.Split(pos[0], "=")[1]
	ys := strings.Split(pos[1], "=")[1]
	zs := strings.Split(strings.Split(pos[2], "=")[1], ">")[0]
	m := new(moon)
	m.x, _ = strconv.Atoi(xs)
	m.y, _ = strconv.Atoi(ys)
	m.z, _ = strconv.Atoi(zs)
	return m
}

func main() {
	silver()
}

func silver() {
	s := system{}
	lines := util.ReadFile("./src/12/input.txt")
	for _, l := range lines {
		s.add(parse(l))
	}
	s.run(1000)
	fmt.Println(s.energy())
}
