package main
import (
	"fmt"
	"sort"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

type parsedInput struct {
	before string
	after  string
}

func parseLine(input string) (parsed parsedInput) {
	parsed.before = input[5:6]
	parsed.after = input[36:37]
	return parsed
}
func depsFor(inputs []parsedInput, step string) (steps []string) {
	steps = []string{}
	for _, input := range inputs {
		if input.after == step {
			steps = append(steps, input.before)
		}
	}
	return steps
}
func dependentsOf(inputs []parsedInput, step string) (steps []string) {
	steps = []string{}
	for _, input := range inputs {
		if input.before == step {
			steps = append(steps, input.after)
		}
	}
	sortSteps(steps)
	return steps
}
func uniqueSteps(inputs []parsedInput) (steps []string) {
	uniqueNames := make(map[string]bool)
	for _, step := range inputs {
		uniqueNames[step.after] = true
		uniqueNames[step.before] = true
	}
	for k := range uniqueNames {
		steps = append(steps, k)
	}
	sortSteps(steps)
	return steps

}
func uniques(stepsWitDupes []string) (steps []string) {
	steps = []string{}
	uniqueNames := make(map[string]bool)
	for _, step := range stepsWitDupes {
		uniqueNames[step] = true
	}
	for k := range uniqueNames {
		steps = append(steps, k)
	}
	sortSteps(steps)
	return steps

}
func sortSteps(steps []string) {
	sort.Slice(steps, func(i, j int) bool {
		return steps[i] < steps[j]
	})
}
func zeroPreqsSteps(inputs []parsedInput) (steps []string) {
	for _, step := range uniqueSteps(inputs) {
		if len(depsFor(inputs, step)) == 0 {
			steps = append(steps, step)
		}
	}
	return steps
}

func availableDependents(inputs []parsedInput, dependents, completedSteps []string) []string {
	available := []string{}
	uniqueSteps := make(map[string]bool)
	for _, step := range completedSteps {
		uniqueSteps[step] = true
	}
	for _, step := range dependents {
		deps := depsFor(inputs, step)
		depsCompleted := true
		for _, dep := range deps {
			if _, ok := uniqueSteps[dep]; !ok {
				depsCompleted = false
				fmt.Println("deps not completed", step, completedSteps, uniqueSteps)
			}
		}
		if depsCompleted {
			available = append(available, step)
		}
	}
	return available
}
func stepOrder(inputs []parsedInput) (stepOrder []string) {
	stepOrder = []string{}
	availableSteps := zeroPreqsSteps(inputs)
	for len(availableSteps) > 0 {
		fmt.Println("step order", stepOrder, "available steps", availableSteps)
		currentStep := availableSteps[0]
		stepOrder = append(stepOrder, currentStep)
		availableSteps = availableSteps[1:]
		dependents := dependentsOf(inputs, currentStep)
		availableDependents := availableDependents(inputs, dependents, stepOrder)
		availableSteps = append(availableSteps, availableDependents...)
		availableSteps = uniques(availableSteps)
		sortSteps(availableSteps)
	}
	return stepOrder
}

func main() {
	// Silver
	lines := util.ReadFile("./day07/input")
	steps := []parsedInput{}
	for _, line := range lines {
		steps = append(steps, parseLine(line))
	}
	sum := stepOrder(steps)
	fmt.Println(strings.Join(sum, ""))
}

