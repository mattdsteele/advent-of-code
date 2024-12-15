package main

import (
	"fmt"
	"strconv"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	// gold()
}

func silver() {
	lines := util.ReadFile("src/4/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	s := Search{}
	s.fields = input
	return strconv.Itoa(s.silver())
}

func gold() {
	lines := util.ReadFile("src/4/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	return "input"
}

type Search struct {
	fields []string
}

type op func(int) int

func nullOp(x int) int {
	return x
}
func incOp(x int) int {
	return x + 1
}
func decOp(x int) int {
	return x - 1
}
func (s Search) checkOp(x, y int, xOp, yOp op) bool {
	return s.letterAt(x, y, "X") &&
		s.letterAt(xOp(x), yOp(y), "M") &&
		s.letterAt(xOp(xOp(x)), yOp(yOp(y)), "A") &&
		s.letterAt(xOp(xOp(xOp(x))), yOp(yOp(yOp(y))), "S")

}
func (s *Search) silver() (count int) {
	// Your logic here to calculate the 'silver' value
	yLen := len(s.fields[0])
	for x := range s.fields {
		y := 0
		for y < yLen {
			count += s.check(x, y)
			y++
		}
	}
	return count
}

func (s Search) check(x, y int) (count int) {
	xActions := []op{nullOp, incOp, decOp}
	yActions := []op{nullOp, incOp, decOp}
	for _, xo := range xActions {
		for _, yo := range yActions {
			if s.checkOp(x, y, xo, yo) == true {
				count++
			}
		}
	}
	return count
}
func (s Search) letterAt(x, y int, letter string) bool {
	if x < 0 {
		return false
	}
	if y < 0 {
		return false
	}
	if x >= len(s.fields) {
		return false
	}
	f := s.fields[x]
	if y >= len(f) {
		return false
	}
	return string(f[y]) == letter
}
