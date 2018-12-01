package main
import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func toSlice(input string) []string {
	return strings.Split(input, "\n")
}
func add(input []string) int {
	sum := 0
	for _, val := range toIArr(input) {
		sum += val
	}
	return sum
}
func firstFrequency(input []string) int {
	intInputs := toIArr(input)
	foundVals := make(map[int]bool)
	runningSum := 0
	foundVals[runningSum] = true
	for {
		for _, val := range intInputs {
			runningSum += val
			_, hasSum := foundVals[runningSum]
			if hasSum {
				return runningSum
			}
			foundVals[runningSum] = true
		}
	}
}

func toIArr(input []string) []int {
	vals := []int{}
	for _, val := range input {
		v, _ := strconv.Atoi(val)
		vals = append(vals, v)
	}
	return vals
}

func main() {
	lines := util.ReadFile("./day01/input.txt")
	firstFreq := firstFrequency(lines)
	fmt.Println(firstFreq)
}

