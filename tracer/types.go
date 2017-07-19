package tracer

type Ray struct {
	Origin    Vector
	Direction Vector
}

func (r *Ray) Point(t float64) Vector {
	return r.Origin.Add(r.Direction.Multiply(t))
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
	Light   MaterialType = -1
	Diffuse              = 0
)
