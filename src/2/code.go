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

func fixAlarm(input []int, firstVal, secondVal int) []int {
	input[1] = firstVal
	input[2] = secondVal
	return input
}

func main() {
	silver()
	gold()
}
func silver() {
	in := util.ReadFile("./src/2/input.txt")[0]
	fmt.Println(opcode(fixAlarm(parse(in), 12, 2))[0])
}

func gold() {
	original := parse(util.ReadFile("./src/2/input.txt")[0])
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copied := make([]int, len(original))
			copy(copied, original)
			fixture := fixAlarm(copied, noun, verb)
			res := opcode(fixture)
			val := res[0]
			if val == 19690720 {
				caledVal := 100*noun + verb
				fmt.Println(caledVal)
				return
			}
		}
	}
}
