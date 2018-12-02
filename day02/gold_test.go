package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestCorrectBoxes(t *testing.T) {
	boxen := `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

	box1, box2 := correctBoxes(util.SliceAtLine(boxen))
	util.Equals(t, box1, "fghij")
	util.Equals(t, box2, "fguij")
}

func TestSimilarItems(t *testing.T) {
	similarLetters := findSimilarLetters("fghij", "fguij")
	util.Equals(t, similarLetters, "fgij")
}

