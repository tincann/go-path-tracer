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

	// scene.Background = Color{R: 55 / 255.0, G: 55 / 255.0, B: 55 / 255.0}

	// lightmat := Material{
	// 	Color: Color{R: 2, G: 2, B: 2},
	// 	Type:  Light,
	// }

	//light
	// scene.AddObject(NewQuad(
	// 	NewVector(-3, 1.5, 2),
	// 	NewVector(-3, 4, 2),
	// 	NewVector(3, 4, 2),
	// 	NewVector(3, 1.5, 2),
	// 	lightmat,
	// ))

	//triangle
	scene.AddObject(NewTriangle(
		NewVector(-0.5, 1, -0.5),
		NewVector(-1.5, 1, -0.5),
		NewVector(-1, 1.5, 0.2),
		Material{
			Color: Color{G: 0.8},
			Type:  Diffuse,
		},
	))

	//sphere
	scene.AddObject(NewSphere(
		NewVector(0, 1.5, -0.5),
		0.5,
		Material{
			Color: Color{R: 0.8},
			Type:  Diffuse,
		},
	))

	scene.AddObject(NewQuad(
		NewVector(5, -2, -1),
		NewVector(5, -2, 1),
		NewVector(5, 2, 1),
		NewVector(5, 2, -1),
		newLightMaterial(Color{1, 1, 1}, 10),
	))

	//floor
	scene.AddObject(NewPlane(
		NewVector(0, 0, 1),
		-1,
		Material{
			Color: Color{R: 0.8, G: 0.8, B: 0.8},
			Type:  Diffuse,
		},
	))

	return scene
}

func newLightMaterial(color Color, intensity float32) Material {
	return Material{
		Color: color.Multiply(intensity),
		Type:  Light,
	}
}
