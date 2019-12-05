package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type rule interface {
	passes(input string) bool
}

type adjacentEqual struct{}

func (a adjacentEqual) passes(input string) bool {
	for i, v := range input {
		if i != 0 {
			prev := input[i-1]
			if v == rune(prev) {
				return true
			}
		}

	}
	return false
}

type digitsDontDecrease struct{}

type concatRules struct {
	rules []rule
}

func initSilverRules() rule {
	s := new(concatRules)
	s.rules = []rule{new(adjacentEqual), new(digitsDontDecrease)}
	return s
}

func initGoldRules() rule {
	s := new(concatRules)
	s.rules = []rule{new(digitsDontDecrease), new(exactlyTwoRepeatingDigits), new(adjacentEqual)}
	return s
}

func (s concatRules) passes(input string) bool {
	for _, r := range s.rules {
		if !r.passes(input) {
			return false
		}
	}
	return true
}

func (d digitsDontDecrease) passes(input string) bool {
	for i, v := range input {
		if i != 0 {
			prev := input[i-1]
			prevVal, _ := strconv.Atoi(string(prev))
			val, _ := strconv.Atoi(string(v))
			if val < prevVal {
				return false
			}
		}
	}
	return true
}

type exactlyTwoRepeatingDigits struct{}

func prevIsSame(i int, input string) bool {
	if i == 0 {
		return false
	}
	prev := rune(input[i-1])
	curr := rune(input[i])
	return prev == curr
}
func nextIsSame(i int, input string) bool {
	if i == len(input)-1 {
		return false
	}
	next := rune(input[i+1])
	curr := rune(input[i])
	return next == curr
}
func followingIsSame(i int, input string) bool {
	if i >= len(input)-2 {
		return false
	}
	following := rune(input[i+2])
	curr := rune(input[i])
	return following == curr
}
func (e exactlyTwoRepeatingDigits) passes(input string) bool {
	for i, _ := range input {
		if !prevIsSame(i, input) {
			if nextIsSame(i, input) {
				if !followingIsSame(i, input) {
					return true
				}
			}
		}
	}
	return false
}

func parse(input string) (string, string) {
	splits := strings.Split(input, "-")
	return splits[0], splits[1]
}
func passingCount(rules rule, input string) int {
	low, high := parse(input)
	passingPasswords := 0
	i, _ := strconv.Atoi(low)
	highVal, _ := strconv.Atoi(high)
	for i < highVal {
		if rules.passes(strconv.Itoa(i)) {
			passingPasswords++
		}
		i++
	}
	return passingPasswords

}
func passingGoldPasswords(input string) int {
	rules := initGoldRules()
	return passingCount(rules, input)
}
func passingSilverPasswords(input string) int {
	rules := initSilverRules()
	return passingCount(rules, input)
}
func main() {
	silver()
	gold()
}
func gold() {
	input := util.ReadFile("src/4/input.txt")[0]
	fmt.Println(passingGoldPasswords(input))

}
func silver() {
	input := util.ReadFile("src/4/input.txt")[0]
	fmt.Println(passingSilverPasswords(input))
}
