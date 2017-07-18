package tracer

import (
	"image/color"
)

type Ray struct {
	Origin    Vector
	Direction Vector
}

type Intersectable interface {
	Intersect(ray Ray) (intersected bool, t float64, n Vector)
	Material() Material
}

type Material struct {
	Color color.Color
	Type  MaterialType
}

type MaterialType int

const (
	Light   MaterialType = 0
	Diffuse              = 1
)
