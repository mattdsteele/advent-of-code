package main
import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	util "github.com/mattdsteele/advent-of-code"
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
		l.action = "wakes up"
	} else {
		l.action = "falls asleep"
	}
}
func uniqueGuards(days []shiftStart) (guardNumbers []string) {
	foundVals := make(map[string]bool)
	for _, day := range days {
		foundVals[day.guard] = true
	}
	for k := range foundVals {
		guardNumbers = append(guardNumbers, k)
	}
	return guardNumbers
}

func sleepiestGuard(shifts []shiftStart, logs []log) (sleepyGuard string) {
	mostMinutesSlept := 0
	for _, guard := range uniqueGuards(shifts) {
		sleptMinutes := totalMinutesAsleep(guard, shifts, logs)
		if sleptMinutes > mostMinutesSlept {
			mostMinutesSlept = sleptMinutes
			sleepyGuard = guard
		}
	}
	return sleepyGuard
}
func sleepiestMinute(guard string, shifts []shiftStart, logs []log) int {
	sleepMap := makeSleepMap(guard, shifts, logs)
	mostSleep := 0
	var sleepiestMinute int
	for i, v := range sleepMap {
		if v > mostSleep {
			mostSleep = v
			sleepiestMinute = i
		}
	}
	return sleepiestMinute
}
func makeSleepMap(guard string, shifts []shiftStart, logs []log) map[int]int {
	sleepHourMinutes := make(map[int]int)
	for _, shift := range shifts {
		if shift.guard == guard {
			dayLogs := logsForDay(logs, shift.startDate)
			logIdx := 0
			asleep := false
			if len(dayLogs) == 0 {
				break
			}
			for i := 0; i < 60; i++ {
				if dayLogs[logIdx].date.Minute() == i {
					if dayLogs[logIdx].action == "falls asleep" {
						asleep = true
					} else {
						asleep = false
					}
					logIdx++
				}
				if asleep {
					sleepHourMinutes[i]++
				}
				if logIdx >= len(dayLogs) {
					break
				}
			}
		}
	}
	return sleepHourMinutes
}
func totalMinutesAsleep(guard string, shifts []shiftStart, logs []log) (minutesAsleep int) {
	for _, v := range makeSleepMap(guard, shifts, logs) {
		minutesAsleep += v
	}
	return minutesAsleep
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
	reg := regexp.MustCompile("^\\d*")
	num := reg.FindString(guardNumber)
	l.guard = num

}
func parseStartShift(input string) shiftStart {
	line := shiftStart{input: input}
	line.parseGuardDate()
	line.parseGuardNumber()
	return line
}
func isShiftStart(input string) bool {
	return strings.Contains(input, "begins shift")
}
func parseFile(input []string) (days []shiftStart, logs []log) {
	for _, entry := range input {
		if isShiftStart(entry) {
			days = append(days, parseStartShift(entry))
		} else {
			logs = append(logs, parseLine(entry))
		}
	}
	sort.Slice(days, func(i, j int) bool { return days[i].startDate.Before(days[j].startDate) })
	sort.Slice(logs, func(i, j int) bool { return logs[i].date.Before(logs[j].date) })
	return days, logs
}

func sameDay(a time.Time, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}
func logsForDay(logs []log, day time.Time) (dayLogs []log) {
	for _, log := range logs {
		if sameDay(log.date, day) {
			dayLogs = append(dayLogs, log)
		}
	}
	return dayLogs
}

func main() {
	lines := util.ReadFile("./day04/input")
	days, logs := parseFile(lines)
	guard := sleepiestGuard(days, logs)
	minute := sleepiestMinute(guard, days, logs)
	guardNo, _ := strconv.Atoi(guard)
	fmt.Println(minute * guardNo)
}

