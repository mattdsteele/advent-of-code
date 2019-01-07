package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"

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

func sortSteps(steps []string) {
	sort.Slice(steps, func(i, j int) bool {
		return steps[i] < steps[j]
	})
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func isAvailable(inputs []parsedInput, step string, completedSteps []string) bool {
	for _, dep := range depsFor(inputs, step) {
		if !contains(completedSteps, dep) {
			return false
		}
	}
	return true
}

func newlyAvailableSteps(inputs []parsedInput, completedSteps, unstartedSteps []string) (newSteps []string) {
	for _, step := range unstartedSteps {
		if isAvailable(inputs, step, completedSteps) {
			newSteps = append(newSteps, step)
		}
	}
	sortSteps(newSteps)
	return newSteps
}

func remove(steps []string, stepToRemove string) (newSteps []string) {
	for _, step := range steps {
		if step != stepToRemove {
			newSteps = append(newSteps, step)
		}
	}
	return newSteps
}

func stepOrder(inputs []parsedInput) (completedSteps []string) {
	unstartedSteps := uniqueSteps(inputs)
	for len(unstartedSteps) > 0 {
		availableSteps := newlyAvailableSteps(inputs, completedSteps, unstartedSteps)
		currentStep := availableSteps[0]
		completedSteps = append(completedSteps, currentStep)
		unstartedSteps = remove(unstartedSteps, currentStep)
	}
	return completedSteps
}

func stepTime(letter string, stepTime int) int {
	asciiVal := int(letter[0])
	return asciiVal - 64 + stepTime
}

func worker(id int, results chan<- int, input <-chan string) {
	for in := range input {
		fmt.Println("here is some input", in, id)
		results <- id
	}
	fmt.Println("done with worker!")
}

type Manager struct {
	done           chan<- bool
	wg             *sync.WaitGroup
	workers        []*Worker
	stepTime       int
	inputs         []parsedInput
	completedSteps []string
	unstartedSteps []string
}

func (m *Manager) tick() {
	newSteps := newlyAvailableSteps(m.inputs, m.completedSteps, m.unstartedSteps)
	fmt.Println("Ticking!", m.completedSteps, newSteps, m.unstartedSteps)
	m.wg.Add(len(m.workers))
	for _, worker := range m.workers {
		if worker.needsWork() && len(newSteps) > 0 {
			step := newSteps[0]
			newSteps = newSteps[1:]
			worker.enqueue(step, stepTime(step, m.stepTime)-1)
			m.unstartedSteps = remove(m.unstartedSteps, step)
			fmt.Println("New unstarted steps:", m.unstartedSteps)
		}
		go worker.tick(m.wg)
	}
	fmt.Println("Ticked all")
	m.wg.Wait()
	m.done <- true
}

type Worker struct {
	id           int
	finishedStep chan<- string
	secondsLeft  int
	workingStep  string
}

func (w *Worker) tick(wg *sync.WaitGroup) {
	defer wg.Done()
	if w.workingStep == "" {
		return
	}
	if w.secondsLeft == 0 {
		w.finishedStep <- w.workingStep
		w.workingStep = ""
		return
	}
	w.secondsLeft--
}

func (w *Worker) enqueue(step string, time int) {
	w.workingStep = step
	w.secondsLeft = time
	fmt.Println(w.id, "Enqueueing step", w.workingStep, w.secondsLeft)
}

func (w *Worker) needsWork() bool {
	return w.workingStep == ""
}

func timeWithHelpers(inputs []parsedInput, stepTime int, workerCount int) int {
	// Setup env
	finishedSteps := make(chan string)
	var workers []*Worker
	for i := 0; i < workerCount; i++ {
		worker := Worker{}
		worker.id = i
		worker.finishedStep = finishedSteps

		workers = append(workers, &worker)
	}
	var wg sync.WaitGroup
	done := make(chan bool)
	manager := Manager{done: done, wg: &wg, workers: workers, stepTime: stepTime, inputs: inputs}

	allSteps := uniqueSteps(inputs)
	manager.unstartedSteps = uniqueSteps(inputs)
	manager.completedSteps = []string{}
	// Run through steps
	i := 0
next:
	for ; len(manager.completedSteps) < len(allSteps); i++ {
		go manager.tick()
		for {
			select {
			case <-done:
				fmt.Println("Done with tick ", i)
				continue next
			case step := <-finishedSteps:
				fmt.Println("Done with step ", step)
				manager.completedSteps = append(manager.completedSteps, step)
			}
		}
	}
	fmt.Println(manager.completedSteps)
	return i
}

func main() {
	lines := util.ReadFile("./day07/input")
	steps := []parsedInput{}
	for _, line := range lines {
		steps = append(steps, parseLine(line))
	}
	totalTime := timeWithHelpers(steps, 60, 5)
	fmt.Println("Total time taken", totalTime)
}

func silverMain() {
	// Silver
	lines := util.ReadFile("./day07/input")
	steps := []parsedInput{}
	for _, line := range lines {
		steps = append(steps, parseLine(line))
	}
	sum := stepOrder(steps)
	fmt.Println(strings.Join(sum, ""))
}

