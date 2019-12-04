package aoc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func InputFor(year, day string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/day/"+day+"/input", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	return CallWithCookie(req)
}
func Stars() int {
	req, _ := http.NewRequest("GET", "https://adventofcode.com", nil)
	res, err := CallWithCookie(req)
	if err != nil {
		log.Fatal("Error calling for stars")
		panic("Failed")
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res))
	if err != nil {
		log.Fatal("Error parsing document")
		panic("Failed")
	}
	withStars := doc.Find(".star-count").First().Text()
	stars, _ := strconv.Atoi(strings.Replace(withStars, "*", "", 1))
	return stars
}

func CallWithCookie(req *http.Request) ([]byte, error) {
	cookie := os.Getenv("AOC_SESSION_COOKIE")
	if cookie == "" {
		panic("Did not set AOC_SESSION_COOKIE")
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
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		panic("Non-OK response back from HTTP call")
	}
	return body, err
}

func Submit(year, day, level, answer string) (bool, string) {
	body := url.Values{}
	body.Add("level", level)
	body.Add("answer", answer)
	req, err := http.NewRequest("POST", "https://adventofcode.com/"+year+"/day/"+day+"/answer", strings.NewReader(body.Encode()))
	if err != nil {
		log.Fatal("Failed to generate request")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := CallWithCookie(req)
	if err != nil {
		log.Fatal("Failed to get val")
	}
	correct := strings.Contains(string(res), "That's the right answer!")
	return correct, string(res)
}
