package main
import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func checkEntry(input string) (containsTwo bool, containsThree bool) {
	foundVals := make(map[string]int)
	runes := toRunes(input)
	for _, rune := range runes {
		count, hasRune := foundVals[rune]
		if !hasRune {
			foundVals[rune] = 1
		} else {
			foundVals[rune] = count + 1
		}
	}
	hasTwo := false
	hasThree := false
	for _, count := range foundVals {
		if count == 2 {
			hasTwo = true
		}
		if count == 3 {
			hasThree = true
		}

	}
	return hasTwo, hasThree
}
func getChecksum(lines []string) int {
	twoCount := 0
	threeCount := 0
	for _, line := range lines {
		hasTwo, hasThree := checkEntry(line)
		if hasTwo {
			twoCount++
		}
		if hasThree {
			threeCount++
		}
	}
	return twoCount * threeCount
}
func toRunes(input string) []string {
	return strings.Split(input, "")
}
func runeDifferences(a string, b string) int {
	differentLetters := 0
	aRunes := toRunes(a)
	bRunes := toRunes(b)
	for i, ch := range aRunes {
		if bRunes[i] != ch {
			differentLetters++
		}
	}
	return differentLetters
}
func correctBoxes(lines []string) (box1 string, box2 string) {
	for i, line := range lines {
		restofLines := lines[i+1:]
		for _, testLine := range restofLines {
			if runeDifferences(line, testLine) == 1 {
				return line, testLine
			}
		}
	}
	panic("nope")
}
func findSimilarLetters(a string, b string) string {
	similarLetters := []string{}
	aRunes := toRunes(a)
	bRunes := toRunes(b)
	for i, ch := range aRunes {
		if bRunes[i] == ch {
			similarLetters = append(similarLetters, ch)
		}
	}
	return strings.Join(similarLetters, "")

}
func silver() {
	lines := util.ReadFile("./day02/input.txt")
	fmt.Println(getChecksum(lines))
}
func main() {
	lines := util.ReadFile("./day02/input.txt")
	fmt.Println(findSimilarLetters(correctBoxes(lines)))
}

