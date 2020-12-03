package main

import (
	"github.com/josephburnett/spel/pkg/app"
)

func main() {
	s, err := app.NewSpel([]string{"testing"})
	if err != nil {
		panic(err)
	}
	for {
		s.Render()
		s.WaitClick()
	}
}
