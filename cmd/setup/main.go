package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	year := os.Getenv("AOC_YEAR")
	if len(os.Args) != 2 {
		panic("Did not pass in a day")
	}
	day := os.Args[1]
	if day == "" {
		panic("Did not pass in a day")
	}
	setupDay(day, year)
}

func setupDay(day string, year string) {
	fmt.Printf("Setting up %s %s\n", day, year)
	padded := fmt.Sprintf("%02s", day)
	dayDir := filepath.Join("./", "day"+padded)
	inputPath := filepath.Join(dayDir, "input.txt")
	fmt.Println(inputPath)
	os.MkdirAll(dayDir, os.ModePerm)
	downloadInput(day, year, inputPath)
}

func downloadInput(day string, year, inputPath string) {
	cookie := os.Getenv("AOC_SESSION_COOKIE")
	fmt.Println(cookie)
	req, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/day/"+day+"/input", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	sessionCookie := http.Cookie{Name: "session", Value: cookie}
	req.AddCookie(&sessionCookie)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	fmt.Printf("%s\n", body)
	if resp.StatusCode != 200 {
		panic("Non-OK response back from HTTP call")
	}
	writeToFile(body, inputPath)

}

func writeToFile(input []byte, inputPath string) {
	err := ioutil.WriteFile(inputPath, input, 0644)
	if err != nil {
		panic(err)
	}
}
