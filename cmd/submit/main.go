package main

import (
	"fmt"
	"os"

	"github.com/mattdsteele/advent-of-code/aoc"
)

func main() {
	year := os.Getenv("AOC_YEAR")
	if len(os.Args) != 4 {
		fmt.Println("usage: ./submit [day] [level] [answer]")
		panic("")
	}
	day := os.Args[1]
	if day == "" {
		panic("Did not pass in a day")
	}
	level := os.Args[2]
	if level != "1" && level != "2" {
		panic("Did not pass in silver (1) or gold (2)")
	}
	answer := os.Args[3]
	if answer == "" {
		panic("Did not pass in an aswer")
	}
	correct, body := aoc.Submit(year, day, level, answer)
	if correct {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
		fmt.Println(body)
	}
}
