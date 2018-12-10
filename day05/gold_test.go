package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestRemovePolarity(t *testing.T) {
	fixtures := [][]string{
		[]string{"A", "dbcCCBcCcD"},
		[]string{"B", "daAcCaCAcCcaDA"},
		[]string{"C", "dabAaBAaDA"},
		[]string{"D", "abAcCaCBAcCcaA"},
	}
	for _, fixture := range fixtures {
		afterReactions := removeFromFixture(fixture[0], "dabAcCaCBAcCcaDA")
		util.Equals(t, fixture[1], afterReactions)
	}
}

