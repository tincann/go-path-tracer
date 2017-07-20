package main

import (
	"image"

	"math"

	"sync"

	"time"

	"fmt"

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

	numRegions := 2
	regions := divideIntoRegions(numRegions, numRegions, bounds)
	go handleEvents(w, screen, regions, tracer)

	w.FlushImage(bounds)
	w.Show()

	scene := t.TriangleScene()
	for {
		wg := sync.WaitGroup{}
		wg.Add(len(regions))

		start := time.Now()

		for _, region := range regions {
			go func(r *ScreenRegion) {
				tracer.TraceRegion(bounds, r.Bounds, r.Accumulator, scene, 10)
				r.Accumulator.DrawContents(screen)
				wg.Done()
			}(region)
		}

		wg.Wait()
		w.FlushImage(bounds)

		fmt.Println(time.Since(start).Seconds())
	}
}

type ScreenRegion struct {
	Bounds      image.Rectangle
	Accumulator *t.Accumulator
}

func divideIntoRegions(xParts, yParts int, screen image.Rectangle) []*ScreenRegion {
	regions := make([]*ScreenRegion, xParts*yParts)
	for y := 0; y < yParts; y++ {
		for x := 0; x < xParts; x++ {
			r := ScreenRegion{}

			xWidth := float64(screen.Dx()) / float64(xParts)
			yWidth := float64(screen.Dy()) / float64(yParts)

			r.Bounds.Min.X = int(math.Floor(float64(x) * xWidth))
			r.Bounds.Min.Y = int(math.Floor(float64(y) * yWidth))

			r.Bounds.Max.X = int(math.Floor(float64(x+1) * xWidth))
			r.Bounds.Max.Y = int(math.Floor(float64(y+1) * yWidth))

			r.Accumulator = t.NewAccumulator(r.Bounds)

			i := xParts*y + x
			regions[i] = &r
		}
	}
	return regions
}

func handleEvents(w wde.Window, screen wde.Image, regions []*ScreenRegion, tracer *t.Tracer) {
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
				resetAccumulators(regions)

			}
		}

		if moveVector.X != 0 || moveVector.Y != 0 || moveVector.Z != 0 {
			tracer.Camera.Move(moveVector)
			resetAccumulators(regions)
		}

		if theta != 0 || phi != 0 {
			tracer.Camera.Rotate(phi, theta)
			resetAccumulators(regions)
		}
	}
}

func resetAccumulators(regions []*ScreenRegion) {
	for _, region := range regions {
		region.Accumulator.Reset()
	}
}
