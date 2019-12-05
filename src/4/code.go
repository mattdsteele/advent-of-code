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

type silverRules struct {
	rules []rule
}

func initSilverRules() rule {
	s := new(silverRules)
	s.rules = []rule{new(adjacentEqual), new(digitsDontDecrease)}
	return s
}

func (s silverRules) passes(input string) bool {
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

func parse(input string) (string, string) {
	splits := strings.Split(input, "-")
	return splits[0], splits[1]
}
func passingSilverPasswords(input string) int {
	low, high := parse(input)
	rules := initSilverRules()
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
func main() {
	silver()
}
func silver() {
	input := util.ReadFile("src/4/input.txt")[0]
	fmt.Println(passingSilverPasswords(input))
}
