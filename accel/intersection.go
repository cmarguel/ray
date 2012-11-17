package accel

import (
	"ray/light/spectrum"
	"ray/mmath"
	"ray/shape"
)

type Intersection struct {
	DiffGeom      shape.DifferentialGeometry
	Primitive     Primitive
	WorldToObject mmath.Transform
	ObjectToWorld mmath.Transform
	PrimitiveId   uint
	RayEpsilon    float64
}

func (i Intersection) Le() spectrum.RGBSpectrum {
	return spectrum.NewRGBSpectrum(0.)
}
