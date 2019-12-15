package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type asteroid struct {
	x int
	y int
}

func (a asteroid) bearing(b asteroid) float64 {
	x := b.x - a.x
	y := b.y - a.y
	deg := (math.Atan2(float64(x), float64(y)) / math.Pi * 180) + 180
	if deg > 360 {
		return 360 - deg
	}
	return 360 - deg
}

func (a asteroid) distance(b asteroid) float64 {
	x := b.x - a.x
	y := b.y - a.y
	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
}

func (f field) visible(a asteroid) int {
	bearings := make(map[float64]bool)
	for _, c := range f.asteroids {
		if a != c {
			b := a.bearing(c)
			bearings[b] = true
		}
	}
	return len(bearings)
}

func (f field) bearingMappings(a asteroid) map[float64][]asteroid {
	bearings := make(map[float64][]asteroid)
	for _, c := range f.asteroids {
		if a != c {
			bearing := a.bearing(c)
			bearings[bearing] = append(bearings[bearing], c)
		}
	}
	return bearings
}

func (f field) vaporized(a asteroid, vaporizedNum int) asteroid {
	sortedMappings := sortMappings(f.bearingMappings(a), a)
	count := 0
	for {
		for i := 0; i < len(sortedMappings); i++ {
			mappings := sortedMappings[i]
			if len(mappings) > 0 {
				count++
				if count == vaporizedNum {
					return mappings[0]
				}
				sortedMappings[i] = mappings[1:]
			}
		}
	}
}

func (f field) mostVisible() (largest int, mostVisible asteroid) {
	for _, a := range f.asteroids {
		visible := f.visible(a)
		if visible > largest {
			largest = visible
			mostVisible = a
		}
	}
	return largest, mostVisible
}

type field struct {
	asteroids []asteroid
}

func main() {
	silver()
	gold()
}

func gold() {
	lines := util.ReadFile("./src/10/input.txt")
	field := parseLines(lines)
	_, a := field.mostVisible()
	final := field.vaporized(a, 200)
	fmt.Println(final.x*100 + final.y)

}
func silver() {
	f := util.ReadFile("./src/10/input.txt")
	i, _ := parseLines(f).mostVisible()
	fmt.Println(i)
}

func parseLines(lines []string) *field {
	f := new(field)
	for y, l := range lines {
		for x, a := range strings.Split(l, "") {
			if a == "#" {
				f.asteroids = append(f.asteroids, asteroid{x, y})
			}
		}
	}
	return f
}

func parse(in string) *field {
	lines := util.SliceAtLine(in)
	return parseLines(lines)
}

func sortMappings(in map[float64][]asteroid, comp asteroid) [][]asteroid {
	// first make a bearing float, which we will sort
	bearings := make([]float64, 0)
	for bearing := range in {
		bearings = append(bearings, bearing)
	}
	sort.Float64s(bearings)

	// next sort each of the elements
	sorted := make([][]asteroid, 0)
	for _, bearing := range bearings {
		asteroids := in[bearing]
		sort.Slice(asteroids, func(i, j int) bool {
			ai := asteroids[i]
			aj := asteroids[j]
			return comp.distance(ai) < comp.distance(aj)
		})
		sorted = append(sorted, asteroids)
	}

	return sorted
}
