package tracer

import (
	"image/color"
)

type Scene struct {
	Objects []Intersectable
	Lights  []*LightSource
}

func NewScene() Scene {
	return Scene{
		Objects: make([]Intersectable, 0),
		Lights:  make([]*LightSource, 0),
	}
}

func (s *Scene) AddObject(t Intersectable) {
	s.Objects = append(s.Objects, t)
}

func (s *Scene) AddLight(t *LightSource) {
	s.Lights = append(s.Lights, t)
}

func TriangleScene() Scene {
	scene := NewScene()

	lightmat := Material{
		Color: color.RGBA{R: 255, G: 255, B: 255},
		Type:  Light,
	}
	mat := Material{
		Color: color.RGBA{G: 128},
		Type:  Diffuse,
	}
	scene.AddLight(NewLightSource(
		NewTriangle(
			NewVector(-0.5, -0.5, 2),
			NewVector(0.5, -0.5, 2),
			NewVector(0, 0.5, 2),
			lightmat),
		color.RGBA{R: 255, G: 255, B: 255}))

	scene.AddObject(NewTriangle(
		NewVector(-0.5, 1, -0.5),
		NewVector(0.5, 1, -0.5),
		NewVector(0, 1, 0.5),
		mat,
	))

	// scene.AddObject(&Triangle{
	// 	P1: NewVector(-0.5, -1, -0.5),
	// 	P2: NewVector(0.5, -1, -0.5),
	// 	P3: NewVector(0, -1, 0.5),
	// })

	return scene
}
