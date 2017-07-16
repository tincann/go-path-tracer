package main

import (
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/xgb"
	t "github.com/tincann/go-path-tracer/src/tracer"
)

func main() {
	go start()
	wde.Run()
}

func start() {
	w, _ := wde.NewWindow(500, 500)
	go handleEvents(w)

	w.Show()

	tracer := t.NewTracer(1, 1, 1)
	for {
		trace(w.Screen(), tracer)
		w.FlushImage(w.Screen().Bounds())
	}
}

func trace(screen wde.Image, tracer *t.Tracer) {
	ray := t.Ray{Start: t.NewVector(0, 0, 0), Direction: t.NewVector(0, 0, 1)}

	for x := 0; x < screen.Bounds().Dx(); x++ {
		for y := 0; y < screen.Bounds().Dy(); y++ {
			c := tracer.TraceRay(ray)
			screen.Set(x, y, c)
		}
	}
}

func handleEvents(w wde.Window) {
	for {
		e := <-w.EventChan()
		switch e.(type) {
		case wde.KeyDownEvent:
			event := e.(wde.KeyDownEvent)
			if event.Key == wde.KeyEscape {
				wde.Stop()
			}
		}
	}
}
