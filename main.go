package main

import _ "github.com/skelterjohn/go.wde/xgb"
import (
	"fmt"

	"github.com/skelterjohn/go.wde"
)

func main() {
	go start()
	wde.Run()
}

func start() {
	fmt.Println("Hello")
	w, _ := wde.NewWindow(500, 500)
	go handleEvents(w)
	w.Show()
}

func handleEvents(w wde.Window) {

	for {
		e := <-w.EventChan()
		switch e.(type) {
		case wde.KeyDownEvent:
			event := e.(wde.KeyDownEvent)
			fmt.Println(event.Key)
			if event.Key == wde.KeyEscape {
				wde.Stop()
			}
		}
	}
}
