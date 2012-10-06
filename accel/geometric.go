package accel

import (
	"container/list"
	"ray/geom"
	"ray/mmath"
	"ray/shape"
)

type GeometricPrimitive struct {
	Shape shape.Shape
}

func NewGeometricPrimitive(shape shape.Shape) GeometricPrimitive {
	return GeometricPrimitive{shape}
}

func (p GeometricPrimitive) CanIntersect() bool {
	return false
}

func (p GeometricPrimitive) Intersect(ray *geom.Ray) (Intersection, bool) {
	dg, _, status := p.Shape.Intersect(ray)

	intersect := Intersection{*dg, p, *new(mmath.Transform), *new(mmath.Transform), nextPrimitiveId(), 0.001}

	return intersect, status
}

func (p GeometricPrimitive) Refine(todo *list.List) {
	refinedShapes := list.New()
	p.Shape.Refine(refinedShapes)
	for e := refinedShapes.Front(); e != nil; e = e.Next() {
		todo.PushBack(GeometricPrimitive{e.Value.(shape.Shape)})
	}
}

func (p GeometricPrimitive) WorldBound() geom.BBox {
	return *new(geom.BBox)
}
