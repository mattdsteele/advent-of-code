package main
import (
	"fmt"
	"runtime"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func react(input string) string {
	var prevRune rune
	for i, r := range input {
		if prevRune != 0 && r != prevRune && (r+32 == prevRune || r-32 == prevRune) {
			// eliminate the offending bits
			remainingRune := input[0:i-1] + input[i+1:]
			return react(remainingRune)
		}

		prevRune = r
	}
	return input
}

func removeFromFixture(upperLetter string, input string) string {
	lowerLetter := strings.ToLower(upperLetter)
	return strings.Replace(strings.Replace(input, upperLetter, "", -1), lowerLetter, "", -1)
}

func mainSilver() {
	// Silver
	lines := util.ReadFile("./day05/input")
	afterReaction := react(lines[0])
	fmt.Println(len(afterReaction))
}

type result struct {
	letter string
	length int
}

func worker(id int, jobs <-chan string, results chan<- result, line string) {
	for letter := range jobs {
		fmt.Println("Worker", id, "Testing", letter)
		testWithoutLetter := removeFromFixture(letter, line)
		afterReaction := react(testWithoutLetter)
		results <- result{letter, len(afterReaction)}
	}
	close(results)
}
func main() {
	// Gold
	line := util.ReadFile("./day05/input")[0]
	cpus := runtime.NumCPU()
	jobs := make(chan string, 100)
	results := make(chan result, 100)
	for i := 0; i < cpus; i++ {
		go worker(i, jobs, results, line)
	}
	for _, letter := range "QWERTYUIOPASDFGHJKLZXCVBNM" {
		jobs <- (string(letter))
	}
	close(jobs)
	smallestAfterReaction := 9999999999999
	for result := range results {
		if result.length < smallestAfterReaction {
			smallestAfterReaction = result.length
			fmt.Println("New smallest letter", result.letter, smallestAfterReaction)
		}
	}
	fmt.Println("Smallest letter is", smallestAfterReaction)
}

