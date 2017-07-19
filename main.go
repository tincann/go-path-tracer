package main

import (
	"image"
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
	acc := t.NewAccumulator(w.Screen().Bounds())

	scene := t.TriangleScene()

	eye := t.NewVector(0, -3, 0)
	direction := t.NewVector(0, 1, 0)
	camera := t.NewCamera(
		eye,
		direction,
		t.NewVector(0, 0, 1),
		1.5, //distance to image plane
		3,   //image plane width
		3,   //image plane height
		0.5, //move speed
	)
	tracer := t.NewTracer(camera)
	go handleEvents(w, w.Screen(), acc, tracer)

	w.FlushImage(w.Screen().Bounds())
	w.Show()

	for {
		trace(w.Screen(), acc, tracer, scene)
		w.FlushImage(acc.Bounds)
		acc.NextFrame()
	}
}

func trace(screen wde.Image, acc *t.Accumulator, tracer *t.Tracer, scene *t.Scene) {
	b := screen.Bounds()
	rayInfos := tracer.Camera.GenerateRays(image.Point{b.Max.X, b.Max.Y}, b)

	for _, rayInfo := range rayInfos {
		c := tracer.TraceRay(rayInfo.Ray, scene, 4)
		avg := acc.SetPixel(rayInfo.X, rayInfo.Y, c)
		screen.Set(rayInfo.X, rayInfo.Y, toSystemColor(avg))
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

func handleEvents(w wde.Window, screen wde.Image, acc *t.Accumulator, tracer *t.Tracer) {
	for {
		e := <-w.EventChan()

		moveVector := t.NewVector(0, 0, 0)
		switch e.(type) {

		case wde.KeyDownEvent:
			event := e.(wde.KeyDownEvent)
			switch event.Key {

			case wde.KeyEscape:
				wde.Stop()

			//camera movement
			case wde.KeyW:
				moveVector.Y++
			case wde.KeyS:
				moveVector.Y--
			case wde.KeyA:
				moveVector.X--
			case wde.KeyD:
				moveVector.X++
			case wde.KeySpace:
				moveVector.Z++
			case wde.KeyLeftShift:
				moveVector.Z--
			}
		}

		if moveVector.X != 0 || moveVector.Y != 0 || moveVector.Z != 0 {
			acc.Reset()
			tracer.Camera.Move(moveVector)
		}
	}
}
