package main

import (
	"fmt"

	"github.com/josephburnett/spel/pkg/word"
)

func main() {
	mutants, err := word.MutateTimes("testing", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(mutants)
}
