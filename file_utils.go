package util
import (
	"bufio"
	"log"
	"os"
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

