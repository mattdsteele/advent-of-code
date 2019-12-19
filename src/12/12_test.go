package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	m := parse("<x=-9, y=6, z=-7>")
	tst.Equals(t, m.x, -9)
	tst.Equals(t, m.y, 6)
	tst.Equals(t, m.z, -7)
	tst.Equals(t, m.dx, 0)
	tst.Equals(t, m.dy, 0)
	tst.Equals(t, m.dz, 0)
	m = parse("<x=5, y=5, z=10>")
	tst.Equals(t, m.x, 5)
	tst.Equals(t, m.y, 5)
	tst.Equals(t, m.z, 10)
}

func TestVelocity(t *testing.T) {
	m := moon{1, 1, 1, 2, 3, 4}
	m.velocity()
	tst.Equals(t, 3, m.x)
}

func TestGravity(t *testing.T) {
	m := parse("<x=-9, y=6, z=-7>")
	o := parse("<x=-9, y=5, z=-6>")
	m.gravity(o)
	tst.Equals(t, 0, m.dx)
	tst.Equals(t, -1, m.dy)
	tst.Equals(t, 1, m.dz)
}

func TestTick(t *testing.T) {
	s := system{}
	s.add(parse("<x=-1, y=0, z=2>"))
	s.add(parse("<x=2, y=-10, z=-7>"))
	s.add(parse("<x=4, y=-8, z=8>"))
	s.add(parse("<x=3, y=5, z=-1>"))
	s.tick()
	m0 := s.moons[0]
	tst.Equals(t, 2, m0.x)
	tst.Equals(t, -1, m0.y)
	tst.Equals(t, 1, m0.z)
}

func TestEnergy(t *testing.T) {
	s := system{}
	s.add(parse("<x=-1, y=0, z=2>"))
	s.add(parse("<x=2, y=-10, z=-7>"))
	s.add(parse("<x=4, y=-8, z=8>"))
	s.add(parse("<x=3, y=5, z=-1>"))
	s.run(10)
	m0 := s.moons[0]
	tst.Equals(t, 36, m0.energy())
	tst.Equals(t, 45, s.moons[1].energy())
}

func TestAnotherExample(t *testing.T) {
	s := system{}
	s.add(parse("<x=-8, y=-10, z=0>"))
	s.add(parse("<x=5, y=5, z=10>"))
	s.add(parse("<x=2, y=-7, z=3>"))
	s.add(parse("<x=9, y=-8, z=-3>"))
	s.run(100)
	tst.Equals(t, 1940, s.energy())
}
