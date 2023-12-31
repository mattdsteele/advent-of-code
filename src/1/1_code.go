package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	return calc(input, silverCalibration)
}
func silverCalibration(s string) int {
	var mappings = map[string]rune{
		"1": '1',
		"2": '2',
		"3": '3',
		"4": '4',
		"5": '5',
		"6": '6',
		"7": '7',
		"8": '8',
		"9": '9',
		"0": '0',
	}
	return calibrationValue(s, mappings)
}

func gold() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(goldCalculate(lines))
}

func calc(lines []string, strategy func(string) int) string {
	total := 0
	for _, line := range lines {
		val := strategy((line))
		total += val
	}
	return fmt.Sprint(total)
}

func goldCalculate(lines []string) string {
	return calc(lines, goldCalibration)
}

func calibrationValue(line string, mappings map[string]rune) int {
	firstPosition := -1
	lastPosition := -1
	var firstValue rune
	var lastValue rune
	for key, val := range mappings {
		first := strings.Index(line, key)
		if first != -1 {
			if firstPosition == -1 || first < firstPosition {
				firstPosition = first
				firstValue = val
			}
		}

		last := strings.LastIndex(line, key)
		if last > lastPosition {
			lastPosition = last
			lastValue = val
		}

	}

	val, _ := strconv.Atoi(string(firstValue) + string(lastValue))
	return val
}

func goldCalibration(s string) int {
	var mappings = map[string]rune{
		"1":     '1',
		"2":     '2',
		"3":     '3',
		"4":     '4',
		"5":     '5',
		"6":     '6',
		"7":     '7',
		"8":     '8',
		"9":     '9',
		"0":     '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	return calibrationValue(s, mappings)
}
