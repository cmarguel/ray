package accel

import (
	"container/list"
	"ray/geom"
	"ray/shape"
)

type GeometricPrimitive struct {
	Shape shape.Shape
}

func (p GeometricPrimitive) CanIntersect() bool {
	return false
}

func (p GeometricPrimitive) Intersect(ray geom.Ray) (Intersection, bool) {
	return *new(Intersection), false
}

func (p GeometricPrimitive) Refine(*list.List) {
}

func (p GeometricPrimitive) WorldBound() BBox {
	return *new(BBox)
}
