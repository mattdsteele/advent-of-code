package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/mattdsteele/advent-of-code/aoc"
)

type Shield struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
}

func Shields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	os.Setenv("AOC_SESSION_COOKIE", "53616c7465645f5f7e791b2ea86f758b14ce2d946ddcd09990a591b875cb1855a6498a5135e74fcd84ec5bf7c6798d77")
	stars := aoc.Stars()
	val := Shield{SchemaVersion: 1, Label: "advent-of-code", Message: fmt.Sprintf("%d stars", stars)}
	j, _ := json.Marshal(val)
	fmt.Fprint(w, string(j))
}
