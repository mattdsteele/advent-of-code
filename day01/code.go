package main
import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func toSlice(input string) []string {
	return strings.Split(input, "\n")
}
func add(input []string) int {
	sum := 0
	for _, val := range input {
		i, _ := strconv.Atoi(val)
		sum += i
	}
	return sum
}
func main() {
	lines := util.ReadFile("./silver-input.txt")
	sum := add(lines)
	fmt.Println(sum)
}

