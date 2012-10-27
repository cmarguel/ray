package accel

import (
	"container/list"
	"ray/geom"
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

type Primitive interface {
	WorldBound() geom.BBox
	CanIntersect() bool
	Intersect(*geom.Ray) (Intersection, bool)
	IntersectP(geom.Ray) bool
	Refine(*list.List)
}

var currPrimitiveId uint = 1

func nextPrimitiveId() uint {
	c := currPrimitiveId
	currPrimitiveId++
	return c
}

func FullyRefine(p Primitive, refined *list.List) {
	todo := list.New()
	todo.PushBack(p)
	for todo.Len() > 0 {
		prim := todo.Remove(todo.Back()).(Primitive)
		if prim.CanIntersect() {
			refined.PushBack(prim)
		} else {
			prim.Refine(todo)
		}
	}
}
