package shape

import (
	"container/list"
	"math"
	"ray/geom"
	"ray/mmath"
)

type Triangle struct {
	V1 geom.Vector3
	V2 geom.Vector3
	V3 geom.Vector3
}

func NewTriangle(
	x1, y1, z1,
	x2, y2, z2,
	x3, y3, z3 float64) Triangle {

	v1 := geom.NewVector3(x1, y1, z1)
	v2 := geom.NewVector3(x2, y2, z2)
	v3 := geom.NewVector3(x3, y3, z3)
	return Triangle{v1, v2, v3}
}

// Taken from PBRT. Mostly the same as the other one, but it'll be easier to work with 
// PBRT's other features if I use their structure and naming conventions.
func (tr Triangle) Intersect(ray *geom.Ray) (*DifferentialGeometry, float64, float64, bool) {
	e1 := tr.V2.Minus(tr.V1)
	e2 := tr.V3.Minus(tr.V1)
	rayD := ray.Direction
	s1 := rayD.Cross(e2)
	divisor := s1.Dot(e1)
	if divisor == 0. {
		return nil, math.Inf(1), 0, false
	}
	invDiv := 1. / divisor

	d := ray.Origin.Minus(tr.V1)
	b1 := d.Dot(s1) * invDiv
	if b1 < 0. || b1 > 1. {
		return nil, math.Inf(1), 0, false
	}

	s2 := d.Cross(e1)
	b2 := rayD.Dot(s2) * invDiv
	if b2 < 0. || b1+b2 > 1. {
		return nil, math.Inf(1), 0, false
	}

	t := e2.Dot(s2) * invDiv
	if t < *ray.MinT || t > *ray.MaxT {
		return nil, math.Inf(1), 0, false
	}

	eps := 1e-3 * t

	dg := DifferentialGeometry{ray.At(t), tr}
	return &dg, t, eps, true
}

func (triangle Triangle) Transform(transform mmath.Transform) Triangle {
	triangle.V1 = transform.Apply(triangle.V1)
	triangle.V2 = transform.Apply(triangle.V2)
	triangle.V3 = transform.Apply(triangle.V3)
	return triangle
}

func (t Triangle) WorldBound() geom.BBox {
	return geom.NewBBox(t.V1, t.V2, t.V3)
}

func (t Triangle) Refine(l *list.List) {
	l.PushBack(t)
}

func (tr Triangle) IntersectP(ray geom.Ray) bool {
	e1 := tr.V2.Minus(tr.V1)
	e2 := tr.V3.Minus(tr.V1)
	rayD := ray.Direction
	s1 := rayD.Cross(e2)
	divisor := s1.Dot(e1)
	if divisor == 0. {
		return false
	}
	invDiv := 1. / divisor

	d := ray.Origin.Minus(tr.V1)
	b1 := d.Dot(s1) * invDiv
	if b1 < 0. || b1 > 1. {
		return false
	}

	s2 := d.Cross(e1)
	b2 := rayD.Dot(s2) * invDiv
	if b2 < 0. || b1+b2 > 1. {
		return false
	}

	t := e2.Dot(s2) * invDiv
	if t < *ray.MinT || t > *ray.MaxT {
		return false
	}

	return true
}

func (t Triangle) CanIntersect() bool {
	return true
}
