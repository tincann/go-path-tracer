package tracer

import "image/color"
import "math"
import "math/rand"

type Tracer struct {
	ViewplaneWidth, ViewplaneHeight float64 //width and height of the view plane
	Distance                        float64 //camera distance from view plane
}

func NewTracer(viewplaneWidth, viewplaneHeight, distance float64) *Tracer {
	return &Tracer{ViewplaneWidth: viewplaneWidth, ViewplaneHeight: viewplaneHeight, Distance: distance}
}

func (t *Tracer) TraceRay(ray Ray, scene Scene) color.Color {
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
		return color.RGBA{
			R: uint8(rand.Int31n(64)),
			G: uint8(rand.Int31n(64)),
			B: uint8(rand.Int31n(64)),
		}
	}
	intersection := ray.Origin.Add(ray.Direction.Multiply(rayT))

	mat := intersectedObj.Material()
	switch mat.Type {
	case Diffuse:
		return t.diffuse(scene, intersection, mat, normal)
	case Light:
		return mat.Color
	}

	//calculate reflected ray
	// p := ray.Origin.Add(ray.Direction.Multiply(rayT)) //intersection point
	// for _, light := range scene.Lights {
	// 	shadowRay := light.Subtract(p)
	// }

	// r := ray.Direction.Subtract(normal.Multiply(2 * ray.Direction.Dot(normal))) //d - 2*(d . n)n

	return color.RGBA{R: 128}
}
