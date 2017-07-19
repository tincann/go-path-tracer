package tracer

import "math"

type Tracer struct {
	Camera *Camera
}

func NewTracer(camera *Camera) *Tracer {
	return &Tracer{Camera: camera}
}

func (t *Tracer) TraceRay(ray Ray, scene *Scene, bouncesLeft int) Color {
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
	}
	intersection := ray.Point(rayT).Add(normal.Multiply(0.001))

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
