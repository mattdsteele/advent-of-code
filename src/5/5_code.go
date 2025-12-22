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
	// gold()
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
	lines := util.ReadFile("src/X/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
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
