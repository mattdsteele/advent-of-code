package main

import (
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	f := parseLine("10 ORE => 10 A")
	target := f.target
	req := f.requires
	tst.Equals(t, 10, target.amount)
	tst.Equals(t, "A", target.chemical)

	tst.Equals(t, 1, len(req))
	r0 := f.requires[0]
	tst.Equals(t, 10, r0.amount)
	tst.Equals(t, "ORE", r0.chemical)
}

func TestParseLonger(t *testing.T) {
	f := parseLine("2 AB, 3 BC, 4 CA => 1 FUEL")
	target := f.target
	req := f.requires
	tst.Equals(t, 1, target.amount)
	tst.Equals(t, "FUEL", target.chemical)

	tst.Equals(t, 3, len(req))
	r0 := f.requires[0]
	tst.Equals(t, 2, r0.amount)
	tst.Equals(t, "AB", r0.chemical)
	tst.Equals(t, 3, f.requires[1].amount)
	tst.Equals(t, "BC", f.requires[1].chemical)
	tst.Equals(t, 4, f.requires[2].amount)
	tst.Equals(t, "CA", f.requires[2].chemical)
}

func TestParseReqs(t *testing.T) {
	r := parseReqs(`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`)
	tst.Equals(t, 6, len(r.pairings))
	tst.Equals(t, "FUEL", r.chemical("FUEL").target.chemical)
	tst.Equals(t, 10, r.chemical("A").target.amount)
}

func TestOreCost(t *testing.T) {
	r := parseReqs(`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`)
	tst.Equals(t, 31, r.oreCost())
}

func TestMoreExamples(t *testing.T) {
	r := parseReqs(`9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`)
	tst.Equals(t, 165, r.oreCost())
}

func TestTable(t *testing.T) {
	tests := []struct {
		in  string
		ore int
	}{{`157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`, 13312}, {`2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`, 180697}, {`171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`, 2210736}}
	for _, test := range tests {
		tst.Equals(t, test.ore, parseReqs(test.in).oreCost())
	}
}
func TestDistance(t *testing.T) {
	r := parseReqs(`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`)
	tst.Equals(t, 1, r.distanceFromOre("A"))
	tst.Equals(t, 1, r.distanceFromOre("B"))
	tst.Equals(t, 2, r.distanceFromOre("C"))
	tst.Equals(t, 3, r.distanceFromOre("D"))
	tst.Equals(t, 4, r.distanceFromOre("E"))
	tst.Equals(t, 5, r.distanceFromOre("FUEL"))
}

func TestLevels(t *testing.T) {
	r := parseReqs(`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`)
	tst.Assert(t, tst.SlicesEqual([]string{"A", "B"}, r.chemicalsAtDistance(1)), "slices not equal")
	tst.Assert(t, tst.SlicesEqual([]string{"C"}, r.chemicalsAtDistance(2)), "slices not equal")
	tst.Assert(t, tst.SlicesEqual([]string{"FUEL"}, r.chemicalsAtDistance(5)), "slices not equal")
}

func TestNextHighestLevel(t *testing.T) {
	s := parseReqs(`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`)
	fuel := s.chemical("FUEL").target
	fuelLevel := []*bond{fuel}
	newDeps := s.depsAtLevel(fuelLevel, 4)
	tst.Equals(t, 2, len(newDeps))
	tst.Equals(t, 7, newDeps[0].amount)

	nextHigherLevel := s.depsAtLevel(newDeps, 3)
	tst.Equals(t, 2, len(nextHigherLevel))
	tst.Equals(t, 14, nextHigherLevel[0].amount)
	tst.Equals(t, "A", nextHigherLevel[0].chemical)
	tst.Equals(t, 1, nextHigherLevel[1].amount)
	tst.Equals(t, "D", nextHigherLevel[1].chemical)
}

func TestConvertNextLevel(t *testing.T) {
	s := parseReqs(`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`)
	aBond := bond{10, "A"}
	converted := s.convert(&aBond)
	tst.Equals(t, 1, len(converted))
	tst.Equals(t, 10, converted[0].amount)
	tst.Equals(t, "ORE", converted[0].chemical)

	underMin := s.convert(&bond{1, "A"})
	tst.Equals(t, 1, len(underMin))
	tst.Equals(t, 10, underMin[0].amount)
	tst.Equals(t, "ORE", underMin[0].chemical)

	overMin := s.convert(&bond{11, "A"})
	tst.Equals(t, 1, len(overMin))
	tst.Equals(t, 20, overMin[0].amount)
	tst.Equals(t, "ORE", overMin[0].chemical)

	multiple := s.convert(&bond{1, "E"})
	tst.Equals(t, 2, len(multiple))
	tst.Equals(t, 7, multiple[0].amount)
	tst.Equals(t, "A", multiple[0].chemical)
	tst.Equals(t, 1, multiple[1].amount)
	tst.Equals(t, "D", multiple[1].chemical)
}
