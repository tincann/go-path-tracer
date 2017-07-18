package tracer

type Scene struct {
	Objects    []Intersectable
	Background Color
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

	scene.Background = Color{R: 111 / 255.0, G: 159 / 255.0, B: 237 / 255.0}

	lightmat := Material{
		Color: Color{R: 1, G: 1, B: 1},
		Type:  Light,
	}
	mat := Material{
		Color:       Color{G: 0.5},
		Type:        Diffuse,
		Specularity: 0.9,
	}

	//light
	scene.AddObject(NewTriangle(
		NewVector(0.5, 1, 2),
		NewVector(-0.5, 1, 2),
		NewVector(0, 2, 2),
		lightmat,
	))

	//triangle
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
