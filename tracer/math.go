package tracer

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l}
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func (v Vector) Multiply(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar}
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X +
		v1.Y*v2.Y +
		v1.Z*v2.Z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
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

func (c Color) MultiplyC(c2 Color) Color {
	return Color{
		R: c.R * c2.R,
		G: c.G * c2.G,
		B: c.B * c2.B,
	}
}

func (c Color) Divide(scalar float32) Color {
	return Color{
		R: c.R / scalar,
		G: c.G / scalar,
		B: c.B / scalar,
	}
}

func Clamp(value, min, max float32) float32 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}

	return value
}
