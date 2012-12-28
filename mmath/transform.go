package mmath

import (
	"math"

	"ray/geom"
)

type Transform struct {
	m   Matrix4x4
	inv Matrix4x4
}

func NewTransform() Transform {
	tr := NewMatrix4x4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
	inv := tr

	return Transform{tr, inv}
}

func (t Transform) Translate(dx, dy, dz float64) Transform {
	tr := NewMatrix4x4(
		1, 0, 0, dx,
		0, 1, 0, dy,
		0, 0, 1, dz,
		0, 0, 0, 1,
	)
	inv := NewMatrix4x4(
		1, 0, 0, -dx,
		0, 1, 0, -dy,
		0, 0, 1, -dz,
		0, 0, 0, 1,
	)
	return Transform{tr.Times(t.m), t.m.Times(inv)}
}

func (t Transform) Scale(x, y, z float64) Transform {
	tr := NewMatrix4x4(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
	inv := NewMatrix4x4(
		1./x, 0, 0, 0,
		0, 1./y, 0, 0,
		0, 0, 1./z, 0,
		0, 0, 0, 1,
	)
	return Transform{tr.Times(t.m), t.m.Times(inv)}
}

func (t Transform) RotateX(angle float64) Transform {
	s := math.Sin(angle)
	c := math.Cos(angle)
	tr := NewMatrix4x4(
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	)
	inv := tr.Transpose()
	return Transform{tr.Times(t.m), t.m.Times(inv)}
}

func (t Transform) RotateY(angle float64) Transform {
	s := math.Sin(angle)
	c := math.Cos(angle)
	tr := NewMatrix4x4(
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	)
	inv := tr.Transpose()
	return Transform{tr.Times(t.m), t.m.Times(inv)}
}

func (t Transform) RotateZ(angle float64) Transform {
	s := math.Sin(angle)
	c := math.Cos(angle)
	tr := NewMatrix4x4(
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
	inv := tr.Transpose()
	return Transform{tr.Times(t.m), t.m.Times(inv)}
}

func (t Transform) LookAt(pos, look, up geom.Vector3) Transform {
	dir := look.Minus(pos).Normalized()
	left := up.Normalized().Cross(dir).Normalized()
	newUp := dir.Cross(left)

	tr := NewMatrix4x4(
		left.X, newUp.X, dir.X, pos.X,
		left.Y, newUp.Y, dir.Y, pos.Y,
		left.Z, newUp.Z, dir.Z, pos.Z,
		0, 0, 0, 1,
	)
	m := tr.Times(t.m)
	return Transform{m, m.Inverse()}
}

func dot(v1, v2 []float64) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (t Transform) Apply(v geom.Vector3) geom.Vector3 {
	v2 := []float64{v.X, v.Y, v.Z, 1.}
	return geom.NewVector3(
		dot(t.m.Row(0), v2),
		dot(t.m.Row(1), v2),
		dot(t.m.Row(2), v2),
	)
}

func (t Transform) ApplyToRay(r geom.Ray) geom.Ray {
	r.Origin = t.Apply(r.Origin)
	r.Direction = t.Apply(r.Direction)
	return r
}

func (t Transform) Times(t2 Transform) Transform {
	m := t.m.Times(t2.m)
	inv := t2.m.Times(t.m)
	return Transform{m, inv}
}

func (t Transform) Perspective(fov, n, f float64) Transform {
	persp := NewMatrix4x4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, f/(f-n), -f*n/(f-n),
		0, 0, 1, 0)

	invTan := 1. / math.Tan(DegToRad*fov/2.)
	m := t.Scale(invTan, invTan, 1).m.Times(persp)
	return Transform{m, m.Inverse()}
}

func (t Transform) Inverse() Transform {
	return Transform{t.inv, t.m}
}
