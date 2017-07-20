package main

import (
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
	screen := w.Screen()
	bounds := screen.Bounds()

	camera := t.NewCamera(
		t.NewVector(0, -3, 0), //eye
		t.NewVector(0, 1, 0),  //direction
		t.NewVector(0, 0, 1),  //up
		1.5,                   //distance to image plane
		3,                     //image plane width
		3,                     //image plane height
		0.5,                   //move speed
	)

	tracer := t.NewTracer(camera, 1, 4)

	acc := t.NewAccumulator(bounds)

	go handleEvents(w, screen, acc, tracer)

	w.FlushImage(bounds)
	w.Show()

	scene := t.TriangleScene()
	for {
		tracer.TraceRegion(bounds, bounds, acc, scene, 10)
		acc.DrawContents(screen)
		w.FlushImage(acc.Bounds)
	}
}

func handleEvents(w wde.Window, screen wde.Image, acc *t.Accumulator, tracer *t.Tracer) {
	for {
		e := <-w.EventChan()

		moveVector := t.NewVector(0, 0, 0)
		theta, phi := 0.0, 0.0
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
			case wde.KeyLeftArrow:
				theta--
			case wde.KeyRightArrow:
				theta++

			case wde.KeyTab:
				acc.Reset()

			}
		}

		if moveVector.X != 0 || moveVector.Y != 0 || moveVector.Z != 0 {
			tracer.Camera.Move(moveVector)
			acc.Reset()
		}

		if theta != 0 || phi != 0 {
			tracer.Camera.Rotate(phi, theta)
			acc.Reset()
		}
	}
}
