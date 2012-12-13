package accel

import (
	"container/list"
	"ray/geom"
	"ray/material"
	"ray/mmath"
	"ray/shape"
)

type GeometricPrimitive struct {
	Shape    shape.Shape
	Material material.Material
}

func NewGeometricPrimitive(shape shape.Shape) GeometricPrimitive {
	return GeometricPrimitive{shape, material.New()}
}

func (p GeometricPrimitive) CanIntersect() bool {
	return true
}

func (p GeometricPrimitive) Intersect(ray *geom.Ray) (Intersection, bool) {
	dg, tHit, eps, status := p.Shape.Intersect(ray)

	if !status {
		return *new(Intersection), false
	}
	intersect := Intersection{*dg, p, *new(mmath.Transform), *new(mmath.Transform), nextPrimitiveId(), eps}
	*ray.MaxT = tHit

	return intersect, true
}

func (p GeometricPrimitive) IntersectP(ray geom.Ray) bool {
	return p.Shape.IntersectP(ray)
}

func (p GeometricPrimitive) Refine(todo *list.List) {
	refinedShapes := list.New()
	p.Shape.Refine(refinedShapes)
	for e := refinedShapes.Front(); e != nil; e = e.Next() {
		todo.PushBack(GeometricPrimitive{e.Value.(shape.Shape), p.Material})
	}
}

func (p GeometricPrimitive) WorldBound() geom.BBox {
	return p.Shape.WorldBound()
}
