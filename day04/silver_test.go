package main
import (
	"testing"
	"time"

	util "github.com/mattdsteele/advent-of-code"
)

func parseTime(val string) time.Time {
	date, err := time.Parse(time.RFC3339, val+"Z")
	if err != nil {
		panic(err)
	}
	return date
}

func TestParse(t *testing.T) {
	type testFixture struct {
		input   string
		time    time.Time
		isAwake bool
	}
	t.Log("Parse")
	tests := []testFixture{
		{"[1518-05-27 00:42] falls asleep", parseTime("1518-05-27T00:42:00"), false},
		{"[1518-05-27 00:42] wakes up", parseTime("1518-05-27T00:42:00"), true},
		{"[2121-01-01 00:01] wakes up", parseTime("2121-01-01T00:01:00"), true},
	}
	for _, test := range tests {
		parsedData := parseLine(test.input)
		util.Assert(t, parsedData.date.Equal(test.time), "dates not equal")
		util.Equals(t, parsedData.isAwake, test.isAwake)
	}
}
func TestParseBeginShift(t *testing.T) {
	type testFixture struct {
		input     string
		guard     string
		startTime time.Time
	}
	tests := []testFixture{
		{"[1518-04-11 00:00] Guard #2207 begins shift", "2207", parseTime("1518-04-11T00:00:00")},
		{"[1518-04-11 00:05] Guard #123 begins shift", "123", parseTime("1518-04-11T00:00:00")},
		{"[1518-09-01 23:52] Guard #3187 begins shift", "3187", parseTime("1518-09-02T00:00:00")},
		{"[1518-09-01 23:32] Guard #1 begins shift", "1", parseTime("1518-09-02T00:00:00")},
		{"[1518-09-02 00:29] Guard #1 begins shift", "1", parseTime("1518-09-02T00:00:00")},
	}
	for _, test := range tests {
		parsedData := parseStartShift(test.input)
		util.Equals(t, parsedData.guard, test.guard)
		util.Assert(t, parsedData.startDate.Equal(test.startTime), "dates not equal")
	}
}

