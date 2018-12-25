package main
import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
)

func reqs() []parsedInput {
	return []parsedInput{
		{before: "C", after: "A"},
		{before: "C", after: "F"},
		{before: "A", after: "B"},
		{before: "A", after: "D"},
		{before: "B", after: "E"},
		{before: "D", after: "E"},
		{before: "F", after: "E"},
	}
}

func TestParsed(t *testing.T) {
	fixtures := [][]string{
		[]string{"Step C must be finished before step F can begin.", "C", "F"},
		[]string{"Step A must be finished before step B can begin.", "A", "B"},
	}
	for _, fixture := range fixtures {
		value := parseLine(fixture[0])
		util.Equals(t, value.before, fixture[1])
		util.Equals(t, value.after, fixture[2])
	}
}
func TestNumberOfPrerequisites(t *testing.T) {
	requirements := reqs()
	type dependencies struct {
		step    string
		prereqs []string
	}
	fixtures := []dependencies{
		{step: "C", prereqs: []string{}},
		{step: "A", prereqs: []string{"C"}},
		{step: "F", prereqs: []string{"C"}},
		{step: "E", prereqs: []string{"B", "D", "F"}},
	}
	for _, fixture := range fixtures {
		deps := depsFor(requirements, fixture.step)
		util.Equals(t, deps, fixture.prereqs)
	}
}

func TestDependents(t *testing.T) {
	type dependents struct {
		step      string
		nextSteps []string
	}
	fixtures := []dependents{
		{step: "C", nextSteps: []string{"A", "F"}},
		{step: "A", nextSteps: []string{"B", "D"}},
	}
	for _, fixture := range fixtures {
		dependents := dependentsOf(reqs(), fixture.step)
		util.Equals(t, dependents, fixture.nextSteps)
	}
}

func TestZeroDepSteps(t *testing.T) {
	reqs := []parsedInput{
		{before: "G", after: "A"},
		{before: "C", after: "A"},
		{before: "C", after: "F"},
		{before: "A", after: "B"},
		{before: "A", after: "D"},
		{before: "B", after: "E"},
		{before: "D", after: "E"},
		{before: "F", after: "E"},
	}
	dependents := zeroPreqsSteps(reqs)
	util.Equals(t, dependents, []string{"C", "G"})
}

func TestSilverOrder(t *testing.T) {
	stepOrder := stepOrder(reqs())
	util.Equals(t, stepOrder, []string{"C", "A", "B", "D", "F", "E"})
}

