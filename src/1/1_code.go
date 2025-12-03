package main

import (
	"fmt"
	"strconv"

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
	status := 50
	count := 0
	for _, i := range input {
		z := delta(i)
		status += z
		if status%100 == 0 {
			count++
		}
		if status >= 100 {
			status -= 100
		}
		if status <= 0 {
			status += 100
		}
	}
	return strconv.Itoa(count)
}

func delta(i string) int {
	val, _ := strconv.Atoi(i[1:])
	diff := i[0:1]
	if diff == "L" {
		val *= -1
	}
	return val
}

func gold() {
	lines := util.ReadFile("src/1/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	status := 50
	prevStatus := 50
	count := 0
	for _, i := range input {
		z := delta(i)
		turns := int(z / 100)
		if turns < 0 {
			turns *= -1
		}
		count += turns
		z = z % 100
		status += z
		if status%100 == 0 {
			count++
		}
		if status > 100 {
			if prevStatus < 100 {
				count++
			}
			status -= 100
		}
		if status < 0 {
			if prevStatus > 0 {
				count++
			}
			status += 100
		}
		prevStatus = status
	}
	return strconv.Itoa(count)
}
