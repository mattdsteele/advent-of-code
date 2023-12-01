package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	util "github.com/mattdsteele/advent-of-code"
	"github.com/mattdsteele/advent-of-code/aoc"
)

func main() {
	year := os.Getenv("AOC_YEAR")
	if len(os.Args) != 2 {
		panic("Did not pass in a day")
	}
	day := os.Args[1]
	if day == "" {
		panic("Did not pass in a day")
	}
	setupDay(day, year)
}

func setupDay(day string, year string) {
	fmt.Printf("Setting up %s %s\n", day, year)
	dayDir := filepath.Join("./", "src", day)
	inputPath := filepath.Join(dayDir, "input.txt")
	fmt.Println(inputPath)
	os.MkdirAll(dayDir, os.ModePerm)
	downloadInput(day, year, inputPath)
	copyTemplateFiles(day)
}

func copyTemplateFiles(day string) {
	templateDir := filepath.Join("./", "template")
	destDir := filepath.Join("./", "src", day)
	err := util.Copy(filepath.Join(templateDir, "code.go"), filepath.Join(destDir, fmt.Sprintf("%s_code.go", day)))
	if err != nil {
		log.Fatal(err)
	}
	err = util.Copy(filepath.Join(templateDir, "test.go"), filepath.Join(destDir, fmt.Sprintf("%s_test.go", day)))
	if err != nil {
		log.Fatal(err)
	}
}

func downloadInput(day string, year, inputPath string) {
	body, err := aoc.InputFor(year, day)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	writeToFile(body, inputPath)

}

func writeToFile(input []byte, inputPath string) {
	err := ioutil.WriteFile(inputPath, input, 0644)
	if err != nil {
		panic(err)
	}
}
