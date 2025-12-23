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
	gold()
}

func silver() {
	lines := util.ReadFile("src/5/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	silver := parse(input).silver()
	return fmt.Sprintf("%d", silver)
}

func gold() {
	lines := util.ReadFile("src/5/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	silver := parse(lines).gold()
	return fmt.Sprintf("%d", silver)
}

type Data struct {
	ranges [][]int
	items  []int
}

func determineType(s string) string {
	if strings.Contains(s, "-") {
		return "RANGES"
	}
	matched, _ := regexp.MatchString(`\d*`, s)
	if matched {
		return "NUMBERS"
	}
	return "NULL"
}
func parse(str []string) *Data {
	d := Data{}
	d.ranges = [][]int{}
	for _, s := range str {
		t := determineType(s)
		switch t {
		case "RANGES":
			ranges := strings.Split(s, "-")
			r1, r2 := ranges[0], ranges[1]
			i1, _ := strconv.Atoi(r1)
			i2, _ := strconv.Atoi(r2)
			d.ranges = append(d.ranges, []int{i1, i2})
		case "NUMBERS":
			i, _ := strconv.Atoi(s)
			d.items = append(d.items, i)
		}
	}
	return &d
}

func (d *Data) inRange(i int) bool {
	for _, r := range d.ranges {
		if inRange(r, i) {
			return true
		}
	}
	return false
}

func (d *Data) silver() (total int) {
	for _, i := range d.items {
		if d.inRange(i) {
			total++
		}
	}
	return total
}

func inRange(r []int, test int) bool {
	low, high := r[0], r[1]
	return test >= low && test <= high
}

func dedupe(input [][]int) [][]int {
	allValidRanges := [][]int{}
	for _, i := range input {
		low, high := i[0], i[1]

		// now check valid ranges
		adjusted := false
		for _, r := range allValidRanges {
			rLow, rHigh := r[0], r[1]
			if low > rLow && high < rHigh {
				// totally within range
				adjusted = true
				continue
			}
			if low < rLow && high < rHigh && high > rLow {
				// adjust the low threshold
				adjusted = true
				r[0] = low
			}

			if low > rLow && low < rHigh && high > rHigh {
				// adjust the high threshold
				adjusted = true
				r[1] = high
			}
		}
		if !adjusted {
			allValidRanges = append(allValidRanges, []int{low, high})
		}
	}
	return allValidRanges
}

func (d *Data) gold() int {
	ranges := len(d.ranges)
	newRanges := dedupe(d.ranges)

	for ranges != len(newRanges) {
		ranges = len(newRanges)
		newRanges = dedupe(newRanges)
	}

	// still need to dedupe values
	total := 0
	// fmt.Println("final ranges")
	// fmt.Println(newRanges)
	for _, r := range newRanges {
		total += r[1] - r[0] + 1
	}
	return total
}
