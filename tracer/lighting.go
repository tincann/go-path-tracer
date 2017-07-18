package tracer

import (
	"image/color"
	"math"
	"math/rand"
)

func (t *Tracer) diffuse(scene Scene, intersection Vector, mat Material, normal Vector) color.Color {
	// ray := Ray{
	// 	Origin:    intersection,
	// 	Direction: uniformHemisphereSample(normal),
	// }

	return mat.Color
}

func reflectRay(ray Ray, intersection Vector) {

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

// Vector3 Sample::CosineSampleHemisphere(float u1, float u2)
// {
//     const float r = Sqrt(u1);
//     const float theta = 2 * kPi * u2;

//     const float x = r * Cos(theta);
//     const float y = r * Sin(theta);

//     return Vector3(x, y, Sqrt(Max(0.0f, 1 - u1)));
// }
