package shape

import (
	"container/list"
	// "math"
	"ray/geom"
	"ray/mmath"
)

type Triangle struct {
	V1 geom.Vertex
	V2 geom.Vertex
	V3 geom.Vertex

	Color geom.Color
}

func NewTriangle(
	x1, y1, z1,
	x2, y2, z2,
	x3, y3, z3 float64) Triangle {
	color := geom.Color{255, 255, 255}
	v1 := geom.Vertex{geom.NewVector3(x1, y1, z1), color}
	v2 := geom.Vertex{geom.NewVector3(x2, y2, z2), color}
	v3 := geom.Vertex{geom.NewVector3(x3, y3, z3), color}
	return Triangle{v1, v2, v3, color}
}

// Taken from PBRT. Mostly the same as the other one, but it'll be easier to work with 
// PBRT's other features if I use their structure and naming conventions.
func (tr Triangle) Intersect(ray *geom.Ray) (*DifferentialGeometry, float64, geom.Color, bool) {
	e1 := tr.V2.P.Minus(tr.V1.P)
	e2 := tr.V3.P.Minus(tr.V1.P)
	rayD := ray.Direction.Minus(ray.Origin)
	s1 := rayD.Cross(e2)
	divisor := s1.Dot(e1)
	if divisor == 0. {
		return nil, 0, *new(geom.Color), false
	}
	invDiv := 1. / divisor

	d := ray.Origin.Minus(tr.V1.P)
	b1 := d.Dot(s1) * invDiv
	if b1 < 0. || b1 > 1. {
		return nil, 0, *new(geom.Color), false
	}

	s2 := d.Cross(e1)
	b2 := rayD.Dot(s2) * invDiv
	if b2 < 0. || b1+b2 > 1. {
		return nil, 0, *new(geom.Color), false
	}

	t := e2.Dot(s2) * invDiv
	if t < *ray.MinT || t > *ray.MaxT {
		return nil, 0, *new(geom.Color), false
	}

	dg := DifferentialGeometry{ray.At(t), tr}
	return &dg, t, tr.Color, true
}

func (triangle Triangle) Transform(transform mmath.Transform) Triangle {
	triangle.V1.P = transform.Apply(triangle.V1.P)
	triangle.V2.P = transform.Apply(triangle.V2.P)
	triangle.V3.P = transform.Apply(triangle.V3.P)
	return triangle
}

func (t Triangle) WorldBound() geom.BBox {
	return geom.NewBBox(t.V1.P, t.V2.P, t.V3.P)
}

func (t Triangle) Refine(l *list.List) {
	l.PushBack(t)
}
