package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	shifts := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		shifts = append(shifts, scanner.Text())
	}
	return shifts
}
func SliceAtLine(input string) []string {
	return strings.Split(input, "\n")
}

