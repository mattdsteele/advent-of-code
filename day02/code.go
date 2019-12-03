package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func parse(input string) (output []int) {
	sInput := strings.Split(input, ",")
	for _, in := range sInput {
		num, _ := strconv.Atoi(in)
		output = append(output, num)
	}
	return output
}

func opcode(input []int) []int {
	p := 0
	for input[p] != 99 {
		ins := input[p : p+4]
		op := ins[0]
		l := input[ins[1]]
		r := input[ins[2]]
		target := ins[3]
		if op == 1 {
			input[target] = l + r
		} else if op == 2 {
			input[target] = l * r
		}
		p += 4
	}
	return input
}

func fixAlarm(input []int) []int {
	input[1] = 12
	input[2] = 2
	return input
}

func main() {
	in := util.ReadFile("./day02/input.txt")[0]
	fmt.Println(opcode(fixAlarm(parse(in)))[0])
}
