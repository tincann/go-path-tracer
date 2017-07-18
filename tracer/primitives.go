package tracer

type Primitive struct {
	Intersectable
	material Material
}

func (p Primitive) Material() Material {
	return p.material
}

type Triangle struct {
	Primitive
	P1, P2, P3 Vector
	Normal     Vector
}

func NewTriangle(p1, p2, p3 Vector, material Material) *Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	n := e1.Cross(e2).Normalize()

	t := Triangle{P1: p1, P2: p2, P3: p3, Normal: n}
	t.material = material
	return &t
}

const epsilon = 1e-8

//Intersect Intersects a ray with a triangle using the Möller–Trumbore algorithm
func (tr *Triangle) Intersect(ray Ray) (intersected bool, t float64, n Vector) {
	n = tr.Normal
	t = 0
	e1 := tr.P2.Subtract(tr.P1)
	e2 := tr.P3.Subtract(tr.P1)

	pvec := ray.Direction.Cross(e2)
	det := e1.Dot(pvec)

	//ray is parallel to plane
	if det < epsilon && det > -epsilon {
		return false, t, n
	}

	invDet := 1 / det
	tvec := ray.Origin.Subtract(tr.P1)
	u := tvec.Dot(pvec) * invDet
	if u < 0 || u > 1 {
		return false, t, n
	}

	qvec := tvec.Cross(e1)
	v := ray.Direction.Dot(qvec) * invDet
	if v < 0 || u+v > 1 {
		return false, t, n
	}

	t = e2.Dot(qvec) * invDet
	return true, t, n
}

//Todo Quad primitive
// type Quad struct {
// 	P1, P2, P3, P4 Vector
// 	Normal         Vector
// }

// func NewQuad(p1, p2, p3, p4 Vector) *Quad {
// 	e1 := p2.Subtract(p1)
// 	e2 := p3.Subtract(p1)
// 	n := e1.Cross(e2).Normalize()
// 	return &Quad{P1: p1, P2: p2, P3: p3, P4: p4, Normal: n}
// }

// func (tr *Quad) Intersect(ray Ray) (intersected bool, t float64, n Vector) {

// }
