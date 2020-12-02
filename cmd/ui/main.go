package main

import (
	"github.com/josephburnett/spel/pkg/app"
	"github.com/josephburnett/spel/pkg/view"
)

func main() {
	s, err := app.NewSpel([]string{"testing"})
	if err != nil {
		panic(err)
	}
	err = view.Render(s.Options(), s.Score())
	if err != nil {
		panic(err)
	}
}
