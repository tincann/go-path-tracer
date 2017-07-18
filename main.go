package main

import (
	"image/color"

	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/xgb"
	t "github.com/tincann/go-path-tracer/tracer"
)

func main() {
	go start()
	wde.Run()
}

func start() {
	w, _ := wde.NewWindow(500, 500)
	go handleEvents(w)

	w.FlushImage(w.Screen().Bounds())
	w.Show()

	tracer := t.NewTracer(2, 2, -1)
	acc := t.NewAccumulator(w.Screen().Bounds())

	for {
		trace(w.Screen(), acc, tracer)
		w.FlushImage(acc.Bounds)
		acc.NextFrame()
	}
}

func trace(screen wde.Image, acc *t.Accumulator, tracer *t.Tracer) {
	//ray := t.Ray{Origin: t.NewVector(0, 0, 0), Direction: t.NewVector(0, -1, 0).Normalize()}
	scene := t.TriangleScene()

	//simple axis aligned camera
	eye := t.NewVector(0, tracer.Distance, 0)

	//viewplane definition
	topleft := t.NewVector(-tracer.ViewplaneWidth/2, 0, tracer.ViewplaneHeight/2)
	topright := t.NewVector(tracer.ViewplaneWidth/2, 0, tracer.ViewplaneHeight/2)
	bottomleft := t.NewVector(-tracer.ViewplaneWidth/2, 0, -tracer.ViewplaneHeight/2)
	e1 := topright.Subtract(topleft)
	e2 := bottomleft.Subtract(topleft)

	maxX, maxY := acc.Bounds.Dx(), acc.Bounds.Dy()

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			vx := float64(x) / float64(maxX)
			vy := float64(y) / float64(maxY)

			p := topleft.Add(e1.Multiply(vx)).Add(e2.Multiply(vy))
			direction := p.Subtract(eye).Normalize()

			ray := t.Ray{Origin: eye, Direction: direction}

			c := tracer.TraceRay(ray, scene, 2)
			avg := acc.SetPixel(x, y, c)
			screen.Set(x, y, toSystemColor(avg))
		}
	}
}

func toSystemColor(c t.Color) color.RGBA {
	return color.RGBA{
		R: uint8(clamp(c.R, 0, 1) * 255),
		G: uint8(clamp(c.G, 0, 1) * 255),
		B: uint8(clamp(c.B, 0, 1) * 255),
	}
}

func clamp(value, min, max float32) float32 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}

	return value
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
