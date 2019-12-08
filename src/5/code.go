package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func parse(input string) (output []string) {
	sInput := strings.Split(input, ",")
	for _, in := range sInput {
		output = append(output, in)
	}
	return output
}

type battlebot struct {
	commands []string
	position int
	opcode3  string
	output   []string
}

func (b *battlebot) paramValue(i instruction, paramNumber int, val int) int {
	paramMode := 0 // position
	if len(i.paramModes) > paramNumber {
		paramMode = i.paramModes[paramNumber]
	}
	if paramMode == 0 {
		v, _ := strconv.Atoi(b.commands[val])
		return v
	}
	return val
}

func (b *battlebot) tick() (done bool) {
	p := b.commands[b.position]
	i := parseInstruction(p)
	switch i.opcode {
	case 1:
		{
			ins := b.commands[b.position : b.position+4]
			li, _ := strconv.Atoi(ins[1])
			ri, _ := strconv.Atoi(ins[2])
			target, _ := strconv.Atoi(ins[3])
			l := b.paramValue(i, 0, li)
			r := b.paramValue(i, 1, ri)
			out := strconv.Itoa(l + r)
			b.commands[target] = out
			b.position += 4
		}
	case 2:
		{
			ins := b.commands[b.position : b.position+4]
			li, _ := strconv.Atoi(ins[1])
			ri, _ := strconv.Atoi(ins[2])
			target, _ := strconv.Atoi(ins[3])
			l := b.paramValue(i, 0, li)
			r := b.paramValue(i, 1, ri)
			out := strconv.Itoa(l * r)
			b.commands[target] = out
			b.position += 4
		}
	case 3:
		{
			newPos, _ := strconv.Atoi(b.commands[b.position+1])
			b.commands[newPos] = b.opcode3
			b.position += 2

		}
	case 4:
		{
			newVal, _ := strconv.Atoi(b.commands[b.position+1])
			code := b.paramValue(i, 0, newVal)
			b.output = append(b.output, strconv.Itoa(code))
			b.position += 2
		}
	case 99:
		{
			done = true
		}
	}
	return done
}

type instruction struct {
	opcode     int
	paramModes []int
}

func opcode(i []string) string {
	chars := len(i)
	if chars < 3 {
		return strings.Join(i, "")
	}
	return strings.Join(i[len(i)-2:], "")
}
func opcodeDigits(i []string) (instruction string, rest []string) {
	o := opcode(i)
	chars := len(i)
	if chars > 2 {
		rest = i[0 : chars-2]
	}
	return o, rest
}
func parseInstruction(i string) (in instruction) {
	is := strings.Split(i, "")
	o, rest := opcodeDigits(is)
	in.opcode, _ = strconv.Atoi(o)
	j := len(rest) - 1
	for j >= 0 {
		p, _ := strconv.Atoi(rest[j])
		in.paramModes = append(in.paramModes, p)
		j--
	}
	return in
}

func make(input, opcode3 string) *battlebot {
	b := new(battlebot)
	b.commands = parse(input)
	b.opcode3 = opcode3
	b.position = 0
	return b
}

func fixAlarm(input []int, firstVal, secondVal int) []int {
	input[1] = firstVal
	input[2] = secondVal
	return input
}

func main() {
	silver()
}
func silver() {
	in := util.ReadFile("./src/5/input.txt")[0]
	bot := make(in, "1")
	for !bot.tick() {
	}
	fmt.Println(bot.output[len(bot.output)-1])
}
