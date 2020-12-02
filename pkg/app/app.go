package app

import (
	"math/rand"
	"sync"

	"github.com/josephburnett/spel/pkg/word"
)

type Spel struct {
	mux         sync.Mutex
	words       []string
	currentWord string
	options     []string
	score       int
}

func NewSpel(words []string) (*Spel, error) {
	s := &Spel{
		words: words,
	}
	s.nextWord()
	return s, nil
}

func (s *Spel) ClickFn(word string) func() {
	return func() {
		s.mux.Lock()
		defer s.mux.Unlock()
		if word == s.currentWord {
			s.score += 3
			s.nextWord()
		} else {
			s.score -= 1
		}
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

func (s *Spel) Options() []string {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.options
}

func (s *Spel) Score() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.score
}
