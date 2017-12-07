package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	var fixtures = []struct {
		input    string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}
	for _, f := range fixtures {
		expected := sequenceSum(f.input)
		if expected != f.expected {
			t.Errorf("Failed with %s, got %d", f.input, f.expected)
		}
	}
	input := "1122"
	expected := sequenceSum(input)
	if expected != 3 {
	}
}

func TestGold(t *testing.T) {
	var fixtures = []struct {
		input    string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}
	for _, f := range fixtures {
		expected := goldOption(f.input)
		if expected != f.expected {
			t.Errorf("Failed with %s, got %d", f.input, f.expected)
		}
	}
	input := "1122"
	expected := sequenceSum(input)
	if expected != 3 {
	}
}
