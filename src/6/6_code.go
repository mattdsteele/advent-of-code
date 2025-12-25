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
	lines := util.ReadFile("src/6/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	return fmt.Sprintf("%d", parse(input).silver())
}

func gold() {
	lines := util.ReadFile("src/6/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(input []string) string {
	return fmt.Sprintf("%d", goldParse(input).gold())
}

type Math struct {
	operations []string
	data       [][]int
	lines      []string
	opLine     string
}

func goldParse(lines []string) *Math {
	math := Math{}
	math.operations = []string{}
	for i, l := range lines {
		if (i + 1) < len(lines) {
			math.lines = append(math.lines, l)
		} else {
			math.opLine = l
		}
	}
	return &math
}

func (m *Math) gold() int {
	total := 0
	i := len(m.opLine) - 1
	data := []string{}
	for i >= 0 {
		val := ""
		for _, l := range m.lines {
			c := string(l[i])
			if c != " " {
				val += c
			}
		}
		if val != "" {
			data = append(data, val)
		}

		op := string(m.opLine[i])
		if op != " " {
			operator := getOperator(op)
			total += goldCalc(data, operator)

			// reset values
			data = []string{}
		}
		i--
	}
	return total
}

func goldCalc(data []string, operator operator) int {
	ints := []int{}
	for _, d := range data {
		i, _ := strconv.Atoi(d)
		ints = append(ints, i)
	}
	return operator(ints)
}

func parse(lines []string) *Math {
	math := Math{}
	math.data = [][]int{}
	for i, l := range lines {
		if (i + 1) < len(lines) {
			dataInt := []int{}
			for _, s := range strings.Split(l, " ") {
				if s != "" {
					i, _ := strconv.Atoi(s)
					dataInt = append(dataInt, i)
				}
			}
			math.data = append(math.data, dataInt)
		} else {
			ops := strings.Split(l, " ")
			for _, o := range ops {
				if o != "" {
					math.operations = append(math.operations, o)
				}
			}
		}
	}
	return &math
}

func (m *Math) silver() int {
	total := 0
	for i, o := range m.operations {
		commands := []int{}
		for _, r := range m.data {
			commands = append(commands, r[i])
		}
		total += line(commands, o)

	}
	return total
}

type operator func([]int) int

func add(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}
func multiply(vals []int) int {
	total := 1
	for _, v := range vals {
		total *= v
	}
	return total
}
func getOperator(op string) operator {
	switch op {
	case "*":
		return multiply
	case "+":
		return add
	}
	panic("Not found: " + op)
}

func line(commands []int, o string) int {
	op := getOperator(o)
	return op(commands)
}
