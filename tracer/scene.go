package tracer

type Scene struct {
	Objects []Intersectable
}

func NewScene() Scene {
	return Scene{
		Objects: make([]Intersectable, 0),
	}
}

func (s *Scene) AddObject(t Intersectable) {
	s.Objects = append(s.Objects, t)
}

func TriangleScene() Scene {
	scene := NewScene()

	lightmat := Material{
		Color: Color{R: 255, G: 255, B: 255},
		Type:  Light,
	}
	mat := Material{
		Color:       Color{G: 128},
		Type:        Diffuse,
		Specularity: 0.9,
	}
	scene.AddObject(NewTriangle(
		NewVector(0.5, 1, 2),
		NewVector(-0.5, 1, 2),
		NewVector(0, 2, 2),
		lightmat,
	))
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
