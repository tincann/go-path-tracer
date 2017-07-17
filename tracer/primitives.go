package tracer

type Ray struct {
	Origin    Vector
	Direction Vector
}

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

type Intersectable interface {
	Intersect(ray Ray) (bool, float64)
}

type Triangle struct {
	P1, P2, P3 Vector
}

const epsilon = 1e-8

//Intersect Intersects a ray with a triangle using the Möller–Trumbore algorithm
func (tr *Triangle) Intersect(ray Ray) (bool, float64) {
	e1 := tr.P2.Subtract(tr.P1)
	e2 := tr.P3.Subtract(tr.P1)

	//plane normal
	pvec := ray.Direction.Cross(e2)
	det := e1.Dot(pvec)

	//ray is parallel to plane
	if det < epsilon && det > -epsilon {
		return false, 0
	}

	invDet := 1 / det
	tvec := ray.Origin.Subtract(tr.P1)
	u := tvec.Dot(pvec) * invDet
	if u < 0 || u > 1 {
		return false, 0
	}

	qvec := tvec.Cross(e1)
	v := ray.Direction.Dot(qvec) * invDet
	if v < 0 || u+v > 1 {
		return false, 0
	}

	t := e2.Dot(qvec) * invDet
	return true, t
}
