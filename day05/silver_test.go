package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func TestPolarityDestroyed(t *testing.T) {
	fixtures := [][]string{
		[]string{"aA", ""},
		[]string{"aa", "aa"},
		[]string{"ab", "ab"},
		[]string{"Aa", ""},
		[]string{"Aab", "b"},
		[]string{"bAab", "bb"},
		[]string{"BAab", ""},
		[]string{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	}
	for _, fixture := range fixtures {
		afterReactions := react(fixture[0])
		util.Equals(t, fixture[1], afterReactions)
	}
}

