package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/josephburnett/spel/pkg/app"
)

var defaultWords = []string{
	"thinking",
	"coding",
	"testing",
	"showing",
}

func main() {
	words := getWords()
	s, err := app.NewSpel(words)
	if err != nil {
		panic(err)
	}
	for {
		s.Render()
		s.WaitClick()
	}
}

func getWords() []string {
	resp, err := http.Get("words.txt")
	if err != nil {
		fmt.Printf("error getting words.txt, using defaults: %v", err)
		return defaultWords
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("non ok status code getting words.txt, using defaults: %v", resp.StatusCode)
		return defaultWords
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading words.txt, using defaults: %v", err)
		return defaultWords
	}
	words := []string{}
	for _, s := range strings.Split(string(body), "\n") {
		if s == "" {
			continue
		}
		words = append(words, s)
	}
	return words
}
