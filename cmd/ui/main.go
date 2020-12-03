package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/josephburnett/spel/pkg/app"
)

func main() {
	resp, err := http.Get("words.txt")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	words := []string{}
	for _, s := range strings.Split(string(body), "\n") {
		if s == "" {
			continue
		}
		words = append(words, s)
	}
	fmt.Printf("words: %v\n", words)
	s, err := app.NewSpel(words)
	if err != nil {
		panic(err)
	}
	for {
		s.Render()
		s.WaitClick()
	}
}
