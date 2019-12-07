package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	sys := new(system)
	sys.parseOrbit("COM)B")
	tst.Equals(t, 2, len(sys.objects))
	orbits := sys.objects
	tst.Equals(t, "COM", orbits[0].name)
	tst.Equals(t, orbits[0], orbits[1].orbiting)
}

func TestFind(t *testing.T) {
	sys := new(system)
	sys.parseOrbit("COM)B")
	sys.parseOrbit("B)C")
	tst.Equals(t, 3, len(sys.objects))

	c := sys.find("C")
	tst.Equals(t, sys.find("B"), c.orbiting)
}

func TestOrbitCount(t *testing.T) {
	sys := new(system)
	sys.parseOrbit("COM)B")
	sys.parseOrbit("B)C")
	tst.Equals(t, 2, sys.find("C").orbitCount())
	tst.Equals(t, 1, sys.find("B").orbitCount())
	tst.Equals(t, 0, sys.find("COM").orbitCount())
}

func TestSystemOrbitCount(t *testing.T) {
	sys := new(system)
	sys.parseOrbit("COM)B")
	sys.parseOrbit("B)C")
	tst.Equals(t, 3, sys.orbitCount())
}

func TestExample(t *testing.T) {
	sys := new(system)
	sys.parseOrbit("COM)B")
	sys.parseOrbit("B)C")
	sys.parseOrbit("C)D")
	sys.parseOrbit("D)E")
	sys.parseOrbit("E)F")
	sys.parseOrbit("B)G")
	sys.parseOrbit("G)H")
	sys.parseOrbit("D)I")
	sys.parseOrbit("E)J")
	sys.parseOrbit("J)K")
	sys.parseOrbit("K)L")
	tst.Equals(t, 42, sys.orbitCount())
}

func TestHops(t *testing.T) {
	sys := new(system)
	sys.parseOrbit("COM)B")
	sys.parseOrbit("B)C")
	sys.parseOrbit("C)D")
	sys.parseOrbit("D)E")
	sys.parseOrbit("E)F")
	sys.parseOrbit("B)G")
	sys.parseOrbit("G)H")
	sys.parseOrbit("D)I")
	sys.parseOrbit("E)J")
	sys.parseOrbit("J)K")
	sys.parseOrbit("K)L")
	sys.parseOrbit("K)YOU")
	sys.parseOrbit("I)SAN")
	tst.Equals(t, sys.find("K"), sys.youOrbiting())
	tst.Equals(t, sys.find("I"), sys.sanOrbiting())
	tst.Equals(t, sys.find("D"), sys.common(sys.youOrbiting(), sys.sanOrbiting()))

	tst.Equals(t, 3, sys.hops(sys.youOrbiting(), sys.find("D")))
	tst.Equals(t, 1, sys.hops(sys.sanOrbiting(), sys.find("D")))

	tst.Equals(t, 4, goldTest(sys))
}
