package tracer

import "math"

type Tracer struct {
	ViewplaneWidth, ViewplaneHeight float64 //width and height of the view plane
	Distance                        float64 //camera distance from view plane
}

func NewTracer(viewplaneWidth, viewplaneHeight, distance float64) *Tracer {
	return &Tracer{ViewplaneWidth: viewplaneWidth, ViewplaneHeight: viewplaneHeight, Distance: distance}
}

func (t *Tracer) TraceRay(ray Ray, scene Scene, bouncesLeft int) Color {
	bouncesLeft--
	if bouncesLeft+1 == 0 {
		return Color{}
	}

	rayT := math.MaxFloat64
	var intersectedObj Intersectable
	var normal Vector
	for _, object := range scene.Objects {
		if yes, t, n := object.Intersect(ray); yes && t < rayT {
			rayT = t
			intersectedObj = object
			normal = n
		}
	}

	if intersectedObj == nil {
		return scene.Background
		// 	R: rand.Float32() * 0.3,
		// 	G: rand.Float32() * 0.3,
		// 	B: rand.Float32() * 0.3,
		// }
	}
	intersection := ray.Origin.Add(ray.Direction.Multiply(rayT)).Add(normal.Multiply(0.00001))

	mat := intersectedObj.Material()
	switch mat.Type {
	case Diffuse:
		return t.diffuse(scene, intersection, mat, normal, bouncesLeft)
	case Light:
		return mat.Color
	}

	//calculate reflected ray
	// p := ray.Origin.Add(ray.Direction.Multiply(rayT)) //intersection point
	// for _, light := range scene.Lights {
	// 	shadowRay := light.Subtract(p)
	// }

	// r := ray.Direction.Subtract(normal.Multiply(2 * ray.Direction.Dot(normal))) //d - 2*(d . n)n

	return Color{}
}
