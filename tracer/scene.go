package tracer

type Scene struct {
	Objects []Intersectable
}

func NewScene() Scene {
	return Scene{
		Objects: make([]Intersectable, 0),
	}
}

func (s *Scene) Add(t *Triangle) {
	s.Objects = append(s.Objects, t)
}

func TriangleScene() Scene {
	scene := NewScene()

	scene.Add(&Triangle{
		P1: NewVector(-0.5, 1, -0.5),
		P2: NewVector(0.5, 1, -0.5),
		P3: NewVector(0, 1, 0.5),
	})
	// scene.Add(&Triangle{
	// 	P1: NewVector(0.5, 1, -0.5),
	// 	P2: NewVector(-0.5, 1, -0.5),
	// 	P3: NewVector(0, 1, 0.5),
	// })
	// scene.Add(&Triangle{
	// 	P1: NewVector(-0.5, -1, -0.5),
	// 	P2: NewVector(0.5, -1, -0.5),
	// 	P3: NewVector(0, -1, 0.5),
	// })

	return scene
}
