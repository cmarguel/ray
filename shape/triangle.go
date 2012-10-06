package shape

import (
	"container/list"
	"math"
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

// Adapted from http://www.softsurfer.com/Archive/algorithm_0105/algorithm_0105.htm
func (t Triangle) Intersect(ray *geom.Ray) (geom.Vector3, geom.Color, bool) {
	u := t.V2.P.Minus(t.V1.P)
	v := t.V3.P.Minus(t.V1.P)
	n := u.Cross(v)

	if n.IsZero() {
		return n, t.Color, false // -1
	}
	dir := ray.Direction.Minus(ray.Origin)
	w0 := ray.Origin.Minus(t.V1.P)
	a := -n.Dot(w0)
	b := n.Dot(dir)

	const delta = 0.00000001
	if math.Abs(b) < delta {
		if math.Abs(a) < delta {
			return n, t.Color, false // 2
		} else {
			return n, t.Color, false // 0
		}
	}
	r := a / b
	if r < 0. {
		return n, t.Color, false // 0
	}

	i := ray.Origin.Plus(dir.Times(r))

	uu := u.Dot(u)
	uv := u.Dot(v)
	vv := v.Dot(v)

	w := i.Minus(t.V1.P)
	wu := w.Dot(u)
	wv := w.Dot(v)

	d := uv*uv - uu*vv

	s := (uv*wv - vv*wu) / d
	if s < 0. || s > 1. {
		return i, t.Color, false // 0
	}

	tt := (uv*wu - uu*wv) / d
	if tt < 0. || s+tt > 1. {
		return i, t.Color, false // 0
	}

	return i, t.Color, true // 1
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
