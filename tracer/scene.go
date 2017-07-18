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
		Color: Color{R: 1, G: 1, B: 1},
		Type:  Diffuse,
	}

	//light
	scene.AddObject(NewQuad(
		NewVector(-5, 0, 2),
		NewVector(-5, 2.5, 2),
		NewVector(5, 2.5, 2),
		NewVector(5, 0, 2),
		lightmat,
	))

	//triangle
	scene.AddObject(NewTriangle(
		NewVector(-1.5, 1, -0.5),
		NewVector(-0.5, 1, -0.5),
		NewVector(-1, 1.5, 0.2),
		mat,
	))

	//sphere
	scene.AddObject(NewSphere(
		NewVector(0, 1, -0.5),
		0.5,
		Material{
			Type: Diffuse,
		},
	))

	//floor
	scene.AddObject(NewPlane(
		NewVector(0, 0, 1),
		-1,
		Material{
			Color: Color{R: 1, G: 1, B: 1},
			Type:  Diffuse,
		},
	))

	return scene
}
