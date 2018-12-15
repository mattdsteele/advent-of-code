package main
import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func toNumbers(input string) (numbers []int) {
	for _, el := range strings.Split(input, " ") {
		asNum, _ := strconv.Atoi(el)
		numbers = append(numbers, asNum)
	}
	return numbers
}

func metadata(els []int, idx int) (total int, endMetadata int) {
	numMetadataElements := els[idx+1]
	numChildren := els[idx]
	childrenSum, idx := 0, idx+2
	for ; numChildren != 0; numChildren-- {
		childSum, childEndIdx := metadata(els, idx)
		childrenSum += childSum
		idx = childEndIdx
	}
	endMetadata = idx + numMetadataElements
	metadataElements := els[idx:endMetadata]
	for _, el := range metadataElements {
		total += el
	}
	return total + childrenSum, endMetadata
}

func checksum(input string, startChar int) (total int) {
	sum, _ := metadata(toNumbers(input), startChar)
	return sum
}
func main() {
	// Silver
	lines := util.ReadFile("./day08/input")
	sum := checksum(lines[0], 0)
	fmt.Println(sum)
}

