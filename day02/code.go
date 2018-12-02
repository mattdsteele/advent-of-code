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
func main() {
	lines := util.ReadFile("./day02/input.txt")
	fmt.Println(getChecksum(lines))

}

