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

func goldResult(s []string) int {
	vals := Chunked(s, 3)
	incremented := 0
	var previous int
	for _, level := range vals {
		cur := sum(level)
		if previous != 0 && previous < cur {
			incremented++
		}
		previous = cur
	}
	return incremented
}

func sum(level []int) int {
	var a int
	for _, i := range level {
		a += i
	}
	return a
}

func gold() {
	lines := util.ReadFile("./src/1/input.txt")
	numLines := goldResult(lines)
	fmt.Println(numLines)
}

func Chunked(lines []string, chunkSize int) (chunkedValues [][]int) {
	asInt := toInt(lines)
	for i := range asInt {
		if i+chunkSize <= len(lines) {
			j := 0
			chunk := []int{}
			for j < chunkSize {
				chunk = append(chunk, asInt[i+j])
				j++
			}
			chunkedValues = append(chunkedValues, chunk)
		}
	}
	return chunkedValues
}

func toInt(lines []string) (asInt []int) {
	for _, line := range lines {
		x, _ := strconv.Atoi(line)
		asInt = append(asInt, x)
	}
	return asInt
}

func IncrementalValues(input []string) int {
	incremented := 0
	var previous int
	for _, level := range input {
		cur, _ := strconv.Atoi(level)
		if previous != 0 && previous < cur {
			incremented++
		}
		previous = cur
	}
	return incremented
}

func silver() {
	lines := util.ReadFile("./src/1/input.txt")
	incremented := IncrementalValues(lines)
	fmt.Println(incremented)

}
