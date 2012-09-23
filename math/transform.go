package math

import (
	"math"
)

type Transform struct {
	m Matrix4x4
}

func NewTransform() Transform {
	return Transform{
		NewMatrix4x4(
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		),
	}
}

func (t Transform) Translate(dx, dy, dz float64) Transform {
	tr := NewMatrix4x4(
		1, 0, 0, dx,
		0, 1, 0, dy,
		0, 0, 1, dz,
		0, 0, 0, 1,
	)
	return Transform{tr.Times(t.m)}
}

func (t Transform) Scale(x, y, z float64) Transform {
	tr := NewMatrix4x4(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
	return Transform{tr.Times(t.m)}
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
	return Transform{tr.Times(t.m)}
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
	return Transform{tr.Times(t.m)}
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
	return Transform{tr.Times(t.m)}
}
