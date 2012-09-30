package accel

import (
	"container/list"
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
	WorldBound() geom.BBox
	CanIntersect() bool
	Intersect(geom.Ray) (Intersection, bool)
	Refine(*list.List)
}

func FullyRefine(p Primitive, refined *list.List) {
	todo := list.New()
	todo.PushBack(p)
	for todo.Len() > 0 {
		prim := todo.Back().Value.(Primitive)
		todo.Remove(todo.Back())
		if prim.CanIntersect() {
			refined.PushBack(prim)
		} else {
			prim.Refine(todo)
		}
	}
}
