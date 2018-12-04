package main
import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type claim struct {
	x      int
	y      int
	width  int
	height int
	id     int
}

func parseLine(input string) claim {
	nameRest := strings.Split(input, "@")
	commaRest := strings.Split(nameRest[1], ",")
	colonRest := strings.Split(commaRest[1], ":")
	xRest := strings.Split(colonRest[1], "x")

	claim := claim{}
	x, _ := strconv.Atoi(strings.TrimSpace(commaRest[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(colonRest[0]))
	width, _ := strconv.Atoi(strings.TrimSpace(xRest[0]))
	height, _ := strconv.Atoi(strings.TrimSpace(xRest[1]))
	claim.x = x
	claim.y = y
	claim.width = width
	claim.height = height
	claim.id, _ = strconv.Atoi(strings.TrimSpace(nameRest[0][1:]))
	return claim
}

func getMaxes(lines []claim) (int, int) {
	maxWidth := 0
	maxHeight := 0
	for _, line := range lines {
		maxWidth = max(maxWidth, line.x+line.width)
		maxHeight = max(maxHeight, line.y+line.height)
	}
	return maxWidth, maxHeight

}
func covers(line claim, x int, y int) bool {
	maxX := line.x + line.width
	maxY := line.y + line.height
	inX := x >= line.x && maxX > x
	inY := y >= line.y && maxY > y
	return inX && inY
}
func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func claimsCount(claims []claim, x int, y int) int {
	fitsCount := 0
	for _, claim := range claims {
		if covers(claim, x, y) {
			fitsCount++
		}
	}
	return fitsCount

}
func mainSilver() {
	claims := []claim{}
	for _, line := range util.ReadFile("./day03/input") {
		claims = append(claims, parseLine(line))
	}
	overlappingSquares := 0
	maxX, maxY := getMaxes(claims)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if claimsCount(claims, x, y) > 1 {
				overlappingSquares++
			}
		}
	}
	fmt.Println("Overlapping squares", overlappingSquares)
}
func main() {
	claims := []claim{}
	for _, line := range util.ReadFile("./day03/input") {
		claims = append(claims, parseLine(line))
	}
	for _, claim := range claims {
		if isExclusive(claims, claim) {
			fmt.Println("Exclusive claim is", claim.id)
		}
	}
}

func isExclusive(claims []claim, test claim) bool {
	for x := test.x; x < test.x+test.width; x++ {
		for y := test.y; y < test.y+test.height; y++ {
			if claimsCount(claims, x, y) != 1 {
				return false
			}
		}
	}
	return true
}

