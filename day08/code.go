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

func metadata(els []int, startChar int) (total int, endMetadata int) {
	numMetadataElements := els[startChar+1]
	numChildren := els[startChar]
	currentChildrenSum, currentChildrenEndMetadata := 0, startChar+2
	for ; numChildren != 0; numChildren-- {
		childSum, endChildMetadata := metadata(els, currentChildrenEndMetadata)
		currentChildrenSum += childSum
		currentChildrenEndMetadata = endChildMetadata + 1
	}
	startMetadata := currentChildrenEndMetadata
	endMetadata = startMetadata + numMetadataElements
	metadataElements := els[startMetadata:endMetadata]
	for _, el := range metadataElements {
		total += el
	}
	return total + currentChildrenSum, endMetadata - 1
}

func checksum(input string, startChar int) (total int) {
	els := toNumbers(input)
	sum, _ := metadata(els, startChar)
	return sum
}
func main() {
	// Silver
	lines := util.ReadFile("./day08/input")
	sum := checksum(lines[0], 0)
	fmt.Println(sum)
}

