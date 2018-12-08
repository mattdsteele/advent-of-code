package main
import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type log struct {
	input   string
	action  string
	date    time.Time
	isAwake bool
}

type shiftStart struct {
	input     string
	guard     string
	startDate time.Time
}

func (l *log) parseAwake() {
	isAwake := l.input[19:]
	if isAwake == "wakes up" {
		l.isAwake = true
	}
}
func getDate(raw string) time.Time {
	dt, err := time.Parse("2006-01-02 15:04", raw[1:17])
	if err != nil {
		panic(err)
	}
	return dt

}
func (l *log) parseDate() {
	l.date = getDate(l.input)
}

func (l *shiftStart) parseGuardDate() {
	guardDate := getDate(l.input)
	newGuardDate := guardDate.Round(time.Hour)
	l.startDate = newGuardDate
}
func parseLine(input string) log {
	// [1518-05-27 00:42] falls asleep
	line := log{input: input}
	line.parseDate()
	line.parseAwake()

	return line
}
func (l *shiftStart) parseGuardNumber() {
	split := strings.Split(l.input, "Guard #")
	guardNumber := split[1]
	fmt.Println(guardNumber)
	reg := regexp.MustCompile("^\\d*")
	num := reg.FindString(guardNumber)
	fmt.Println(num)
	l.guard = num

}
func parseStartShift(input string) shiftStart {
	line := shiftStart{input: input}
	line.parseGuardDate()
	line.parseGuardNumber()
	return line
}

func main() {}

