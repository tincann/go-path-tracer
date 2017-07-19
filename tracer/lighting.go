package tracer

import (
	"math"
	"math/rand"
)

func (t *Tracer) diffuse(scene *Scene, intersection Vector, mat Material, normal Vector, bouncesLeft int) Color {
	ray := Ray{
		Origin:    intersection,
		Direction: uniformHemisphereSample(normal),
	}

	brdf := mat.Color.Divide(math.Pi)
	nDotR := normal.Dot(ray.Direction)
	Ei := t.TraceRay(ray, scene, bouncesLeft).Multiply(float32(nDotR))

	return Ei.MultiplyC(brdf).Multiply(math.Pi * 2)
}

func uniformHemisphereSample(orientation Vector) Vector {
	theta := rand.Float64() * 2 * math.Pi
	u := rand.Float64()*2 - 1
	x := math.Sqrt(1 - u*u)
	v := Vector{
		X: x * math.Cos(theta),
		Y: x * math.Sin(theta),
		Z: u,
	}

	if v.Dot(orientation) < 0 {
		return v.Multiply(-1)
	}

	return v
}
