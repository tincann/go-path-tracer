package tracer

import (
	"math"
)

type Primitive struct {
	Intersectable
	material Material
}

func (p Primitive) Material() Material {
	return p.material
}

type Plane struct {
	Primitive
	Normal Vector
	D      float64
}

func NewPlane(normal Vector, d float64, material Material) *Plane {
	p := Plane{
		Normal: normal.Normalize(),
		D:      d,
	}
	p.material = material
	return &p
}

func (p *Plane) Intersect(ray Ray) (intersected bool, t float64, n Vector) {
	t = (ray.Origin.Dot(p.Normal) + p.D) / ray.Direction.Dot(p.Normal)
	return t >= 0, t, p.Normal
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
	// if ray.Direction.Dot(n) < 0 {
	// 	n = n.Multiply(-1)
	// }

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

type Quad struct {
	Primitive
	P1, P2, P3, P4 Vector
	Normal         Vector
}

func NewQuad(p1, p2, p3, p4 Vector, material Material) *Quad {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	n := e1.Cross(e2).Normalize()
	q := Quad{P1: p1, P2: p2, P3: p3, P4: p4, Normal: n}
	q.material = material
	return &q
}

//todo don't use triangles
func (q *Quad) Intersect(ray Ray) (intersected bool, t float64, n Vector) {
	t1 := Triangle{P1: q.P1, P2: q.P2, P3: q.P4}
	t2 := Triangle{P1: q.P2, P2: q.P3, P3: q.P4}

	if intersected, t, n := t1.Intersect(ray); intersected {
		return intersected, t, n
	}

	if intersected, t, n := t2.Intersect(ray); intersected {
		return intersected, t, n
	}

	return false, t, n
}

type Sphere struct {
	Primitive
	Center       Vector
	Radius, rad2 float64
}

func NewSphere(center Vector, radius float64, material Material) *Sphere {
	s := Sphere{Center: center, Radius: radius, rad2: radius * radius}
	s.material = material
	return &s
}

func (s *Sphere) Intersect(ray Ray) (intersected bool, t float64, n Vector) {
	c := ray.Origin.Subtract(s.Center)
	t = c.Multiply(-1).Dot(ray.Direction)
	q := c.Multiply(-1).Subtract(ray.Direction.Multiply(t))
	p2 := q.Dot(q)
	if p2 > s.rad2 {
		return false, t, n
	}
	t -= math.Sqrt(s.rad2 - p2)

	if t > 0 {
		p := ray.Origin.Add(ray.Direction.Multiply(t))
		n = p.Subtract(s.Center).Normalize()
		return true, t, n
	}

	return false, t, n
}

// //fast way to calculate
// t = Vector3.Dot(-C, ray.Direction);
// var q = -C - t*ray.Direction;
// float p2 = Vector3.Dot(q, q);
// if (p2 > _rad2) return false;
// t -= (float) Math.Sqrt(_rad2 - p2);
