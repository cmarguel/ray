package geom

import (
	"math"
)

type Point Vector3
type Color Vector3

type Vertex struct {
	P Vector3
	C Color
}

type Triangle struct {
	V1 Vertex
	V2 Vertex
	V3 Vertex
}

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

// Adapted from http://www.softsurfer.com/Archive/algorithm_0105/algorithm_0105.htm
func (ray Ray) IntersectTriangle(t Triangle) (Vector3, int) {
	u := t.V2.P.Minus(t.V1.P)
	v := t.V3.P.Minus(t.V1.P)
	n := u.Cross(v)

	if n.IsZero() {
		return n, -1
	}
	dir := ray.Direction.Minus(ray.Origin)
	w0 := ray.Origin.Minus(t.V1.P)
	a := -n.Dot(w0)
	b := n.Dot(dir)

	const delta = 0.00000001
	if math.Abs(b) < delta {
		if math.Abs(a) < delta {
			return n, 2
		} else {
			return n, 0
		}
	}
	r := a / b
	if r < 0. {
		return n, 0
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
		return i, 0
	}

	tt := (uv*wu - uu*wv) / d
	if tt < 0. || s+tt > 1. {
		return i, 0
	}

	return i, 1
}
