package main

import (
	"strings"
	"testing"

	tst "github.com/mattdsteele/advent-of-code/testing"
)

func TestParse(t *testing.T) {
	data := "123456789012"
	img := make(3, 2, data)
	tst.Equals(t, 2, len(img.layers))
	tst.Equals(t, strings.Split("123456", ""), img.layers[0].data)
}
func TestCountDigits(t *testing.T) {
	data := "123456789012"
	img := make(3, 2, data)
	tst.Equals(t, 1, img.layers[1].digitsOf("0"))
	tst.Equals(t, 0, img.layers[1].digitsOf("4"))
}

func TestLeastWith(t *testing.T) {
	data := "222233223333"
	img := make(3, 2, data)
	tst.Equals(t, img.layers[1], img.leastWith("2"))
	tst.Equals(t, img.layers[0], img.leastWith("3"))
}
