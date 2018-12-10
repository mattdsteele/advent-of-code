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

func main() {
	lines := util.ReadFile("./day05/input")
	afterReaction := react(lines[0])
	fmt.Println(len(afterReaction))
}

