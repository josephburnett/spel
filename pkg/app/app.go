package app

import (
	"fmt"
	"math/rand"
	"sync"
	"syscall/js"

	"github.com/josephburnett/spel/pkg/word"
)

type Spel struct {
	mux         sync.Mutex
	words       []string
	currentWord string
	options     []string
	score       int
	click       chan struct{}
	newCat      bool
	catWidth    int
}

func NewSpel(words []string) (*Spel, error) {
	s := &Spel{
		words: words,
		click: make(chan struct{}),
	}
	s.nextWord()
	return s, nil
}

func (s *Spel) clickFn(word string) func(js.Value, []js.Value) interface{} {
	return func(_ js.Value, _ []js.Value) interface{} {
		s.mux.Lock()
		defer s.mux.Unlock()
		if word == s.currentWord {
			s.score += 3
			s.nextWord()
		} else {
			s.score -= 1
			if s.score < 0 {
				s.score = 0
			}
			s.catWidth = s.catWidth / 2
		}
		s.click <- struct{}{}
		return nil
	}
}

func (s *Spel) nextWord() error {
	s.currentWord = s.words[rand.Intn(len(s.words))]
	options, err := word.MutateTimes(s.currentWord, 8)
	if err != nil {
		return err
	}
	options = append(options, s.currentWord)
	rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	s.options = options
	s.catWidth = 500
	s.newCat = true
	return nil
}

func (s *Spel) Render() {
	doc := js.Global().Get("document")
	app := doc.Call("getElementById", "app")
	app.Set("style", "float:left;")
	app.Set("innerHTML", "")

	top := doc.Call("createElement", "div")
	top.Set("style", "float:left;clear:both;height:100px")
	title := doc.Call("createElement", "h1")
	title.Set("innerHTML", fmt.Sprintf("Score: %v", s.score))
	top.Call("appendChild", title)
	app.Call("appendChild", top)

	ul := doc.Call("createElement", "ul")
	ul.Set("style", "width:300px;float:left;font-size:2em;clear:left")
	for _, word := range s.options {
		li := doc.Call("createElement", "li")
		li.Set("onclick", js.FuncOf(s.clickFn(word)))
		text := doc.Call("createTextNode", word)
		li.Call("appendChild", text)
		ul.Call("appendChild", li)
	}
	app.Call("appendChild", ul)

	style := fmt.Sprintf("width:%vpx;float:left", s.catWidth)
	if s.newCat {
		cat := doc.Call("getElementById", "cat")
		catImg := doc.Call("createElement", "img")
		catImg.Set("src", fmt.Sprintf("https://cataas.com/cat?fresh=%v", rand.Int()))
		catImg.Set("id", "cat-image")
		catImg.Set("style", style)
		cat.Set("innerHTML", "")
		cat.Call("appendChild", catImg)
		s.newCat = false
	} else {
		catImg := doc.Call("getElementById", "cat-image")
		catImg.Set("style", style)
	}
}

func (s *Spel) WaitClick() {
	<-s.click
}
