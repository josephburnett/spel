package app

import (
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
	return nil
}

func (s *Spel) Render() {
	doc := js.Global().Get("document")
	app := doc.Call("getElementById", "app")
	ul := doc.Call("createElement", "ul")
	for _, word := range s.options {
		li := doc.Call("createElement", "li")
		li.Set("onclick", js.FuncOf(s.clickFn(word)))
		text := doc.Call("createTextNode", word)
		li.Call("appendChild", text)
		ul.Call("appendChild", li)
	}
	app.Set("innerHTML", "")
	app.Call("appendChild", ul)
}

func (s *Spel) WaitClick() {
	<-s.click
}
