package accel

import (
	"ray/geom"
	"ray/mmath"
)

type Intersection struct {
	Primitive     Primitive
	WorldToObject mmath.Transform
	ObjectToWorld mmath.Transform
	PrimitiveId   uint
	RayEpsilon    float64
}

type Primitive interface {
	CanIntersect() bool
	Intersect(geom.Ray) (Intersection, bool)
	Refine([]Primitive)
}
