package util

import (
	"bufio"
	"io"
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

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
