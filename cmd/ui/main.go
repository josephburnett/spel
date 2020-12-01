package main

import (
	"fmt"
	"syscall/js"

	"github.com/josephburnett/spel/pkg/word"
)

func main() {
	doc := js.Global().Get("document")
	app := doc.Call("getElementById", "app")
	mutants, err := word.MutateTimes("testing", 10)
	if err != nil {
		panic(err)
	}
	app.Set("innerHTML", fmt.Sprintf("%v", mutants))
}
