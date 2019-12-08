package main

import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type image struct {
	width  int
	height int
	layers []*layer
}

func (i image) leastWith(digit string) (found *layer) {
	mostDigits := 99999999999999999
	for _, l := range i.layers {
		foundDigits := l.digitsOf(digit)
		if foundDigits < mostDigits {
			mostDigits = foundDigits
			found = l
		}
	}
	return found
}

func make(width, height int, input string) image {
	i := image{width, height, nil}
	layerSize := width * height
	for len(input) > 0 {
		newLayer := input[0:layerSize]
		i.layers = append(i.layers, &layer{strings.Split(newLayer, "")})
		input = input[layerSize:]
	}
	return i
}

type layer struct {
	data []string
}

func (l layer) digitsOf(digit string) (count int) {
	for _, d := range l.data {
		if d == digit {
			count++
		}
	}
	return count
}

func main() {
	silver()
}

func silver() {
	in := util.ReadFile("./src/8/input.txt")[0]
	image := make(25, 6, in)
	least := image.leastWith("0")
	fmt.Println(least.digitsOf("1") * least.digitsOf("2"))
}
