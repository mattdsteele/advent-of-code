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

func sumMetadata(els []int, idx int) (total int, endMetadata int) {
	numMetadataElements := els[idx+1]
	numChildren := els[idx]
	if numChildren != 0 {
		panic("this cannot be non-zero")
	}
	idx += 2
	endMetadata = idx + numMetadataElements
	metadataElements := els[idx:endMetadata]
	for _, el := range metadataElements {
		total += el
	}
	return total, endMetadata
}
func value(els []int, idx int) (total int, endMetadata int) {
	numMetadataElements := els[idx+1]
	numChildren := els[idx]
	if numChildren == 0 {
		return sumMetadata(els, idx)
	}
	idx += 2
	childValues := make(map[int]int)
	for i := 0; i < numChildren; i++ {
		childValue, childEndIdx := value(els, idx)
		childValues[i+1] = childValue
		idx = childEndIdx
	}
	endMetadata = idx + numMetadataElements
	metadataElements := els[idx:endMetadata]
	for _, el := range metadataElements {
		value, found := childValues[el]
		if found {
			total += value
		}
	}
	return total, endMetadata
}

func checksum(input string, startChar int) (total int) {
	sum, _ := metadata(toNumbers(input), startChar)
	return sum
}
func goldValue(input string, startChar int) (total int) {
	sum, _ := value(toNumbers(input), startChar)
	return sum
}
func mainSilver() {
	// Silver
	lines := util.ReadFile("./day08/input")
	sum := checksum(lines[0], 0)
	fmt.Println(sum)
}
func main() {
	// gold
	lines := util.ReadFile("./day08/input")
	sum := goldValue(lines[0], 0)
	fmt.Println(sum)
}

