package view

import "syscall/js"

func Render(options []string, score int) error {
	doc := js.Global().Get("document")
	app := doc.Call("getElementById", "app")
	ul := doc.Call("createElement", "ul")
	app.Call("appendChild", ul)
	for _, word := range options {
		li := doc.Call("createElement", "li")
		ul.Call("appendChild", li)
		text := doc.Call("createTextNode", word)
		li.Call("appendChild", text)
	}
	return nil
}

// s := js.Global.Get("document").Call("createElement", "span")
// f.Call("appendChild", s)
// t := js.Global.Get("document").Call("createTextNode", "Hello World")
// s.Call("appendChild", t)
