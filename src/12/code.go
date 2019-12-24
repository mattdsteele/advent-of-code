package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type moon struct {
	x       int
	y       int
	z       int
	dx      int
	dy      int
	dz      int
	startX  int
	startY  int
	startZ  int
	startDX int
	startDY int
	startDZ int
}

func (m *moon) toString() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d,%d", m.x, m.y, m.z, m.dx, m.dy, m.dz)
}

type system struct {
	moons   []*moon
	repeatX int64
	repeatY int64
	repeatZ int64
}

func (s *system) toString() string {
	st := ""
	for _, m := range s.moons {
		st += m.toString()
	}
	return st
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

func (s *system) setInitial() {
	for _, m := range s.moons {
		m.startX = m.x
		m.startY = m.y
		m.startZ = m.z
		m.startDX = m.dx
		m.startDY = m.dy
		m.startDZ = m.dz
	}
}

func (s *system) run(times int) {
	s.setInitial()
	i := 0
	for i < times {
		s.tick(int64(i))
		i++
	}
}

func (s *system) repeats() int64 {
	s.setInitial()
	i := int64(0)
	for {
		i++
		if s.tick(i) {
			return s.lcm()
		}
	}
}

func (s *system) lcm() int64 {
	l := int64(1)
	l = lcm(l, s.repeatX)
	l = lcm(l, s.repeatY)
	l = lcm(l, s.repeatZ)
	return l
}

func (s *system) tick(step int64) bool {
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
	s.checkLoop(step)
	return s.allLooped()
}
func (s *system) checkLoop(step int64) {
	if s.repeatX == int64(0) && s.checkX() {
		s.repeatX = step
	}
	if s.repeatY == int64(0) && s.checkY() {
		s.repeatY = step
	}
	if s.repeatZ == int64(0) && s.checkZ() {
		s.repeatZ = step
	}

}
func (s *system) checkX() bool {
	for _, m := range s.moons {
		if m.startX != m.x || m.startDX != m.dx {
			return false
		}
	}
	return true
}
func (s *system) checkY() bool {
	for _, m := range s.moons {
		if m.startY != m.y || m.startDY != m.dy {
			return false
		}
	}
	return true
}
func (s *system) checkZ() bool {
	for _, m := range s.moons {
		if m.startZ != m.z || m.startDZ != m.dz {
			return false
		}
	}
	return true
}
func (s *system) allLooped() bool {
	if s.repeatX != int64(0) && s.repeatY != int64(0) && s.repeatZ != int64(0) {
		return true
	}
	return false
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
	gold()
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
func gold() {
	s := system{}
	lines := util.ReadFile("./src/12/input.txt")
	for _, l := range lines {
		s.add(parse(l))
	}
	fmt.Println(s.repeats())
}

// greatest common divisor (gcd) via Euclidean algorithm
func gcd(a, b int64) int64 {
	for b != int64(0) {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}
