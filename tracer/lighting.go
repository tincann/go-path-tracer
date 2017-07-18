package tracer

import (
	"math"
	"math/rand"
)

func (t *Tracer) diffuse(scene Scene, intersection Vector, mat Material, normal Vector, bouncesLeft int) Color {
	ray := Ray{
		Origin:    intersection,
		Direction: uniformHemisphereSample(normal),
	}

	diffuseColor := mat.Color.Multiply(1 - mat.Specularity)
	reflectedColor := t.TraceRay(ray, scene, bouncesLeft).Multiply(mat.Specularity)
	return diffuseColor.Add(reflectedColor)
}

func (c Color) Add(c2 Color) Color {
	return Color{
		R: c.R + c2.R,
		G: c.G + c2.G,
		B: c.B + c2.B,
	}
}

func (c Color) Multiply(scalar float32) Color {
	return Color{
		R: c.R * scalar,
		G: c.G * scalar,
		B: c.B * scalar,
	}
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
