package tracer

type Ray struct {
	Origin    Vector
	Direction Vector
}

type Intersectable interface {
	Intersect(ray Ray) (intersected bool, t float64, n Vector)
	Material() Material
}

type Material struct {
	Color       Color
	Type        MaterialType
	Specularity float32
}

type Color struct {
	R, G, B float32
}

type MaterialType int

const (
	Light   MaterialType = 0
	Diffuse              = 1
)
