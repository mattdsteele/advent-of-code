package main
import (
	"fmt"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func react(input string) string {
	runes := strings.Split(input, "")
	var prevRune string
	for i, r := range runes {
		if prevRune != "" && strings.ToUpper(r) == strings.ToUpper(prevRune) && r != prevRune {
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
func main() {
	// Gold
	line := util.ReadFile("./day05/input")[0]
	var smallestLetter int
	for _, letter := range "QWERTYUIOPASDFGHJKLZXCVBNM" {
		testWithoutLetter := removeFromFixture(string(letter), line)
		fmt.Println("Testing ", string(letter))
		afterReaction := react(testWithoutLetter)
		sizeAfterReaction := len(afterReaction)
		if smallestLetter == 0 || sizeAfterReaction < smallestLetter {
			smallestLetter = sizeAfterReaction
			fmt.Println("New smallest letter", string(letter), sizeAfterReaction)
		}
	}
}

