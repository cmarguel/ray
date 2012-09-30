package accel

import (
	"container/list"
	"ray/geom"
)

type Voxel struct {
	Primitives      *list.List
	AllCanIntersect bool
}

func NewVoxel(p Primitive) Voxel {
	primitives := list.New()
	primitives.PushBack(p)
	return Voxel{primitives, false}
}

func (v Voxel) AddPrimitive(p Primitive) {
	v.Primitives.PushBack(p)
}

func (v Voxel) Intersect(ray *geom.Ray) (Intersection, bool) {
	if !v.AllCanIntersect {
		// TODO write lock here 
		for e := v.Primitives.Front(); e != nil; e = e.Next() {
			prim := e.Value.(Primitive)
			refined := v.refine(prim)
			e.Value = refined
		}
		v.AllCanIntersect = true
		// TODO remove write lock
	}

	return v.findIntersections(ray)
}

func (v Voxel) findIntersections(ray *geom.Ray) (Intersection, bool) {
	hitSomething := false
	var intersection = *new(Intersection)

	for e := v.Primitives.Front(); e != nil; e = e.Next() {
		prim := e.Value.(Primitive)
		found := false
		intersection, found = prim.Intersect(ray)
		if found {
			hitSomething = true
		}
	}
	return intersection, hitSomething
}

func (v Voxel) refine(prim Primitive) Primitive {
	if !prim.CanIntersect() {
		p := list.New()
		FullyRefine(prim, p)
		if p.Len() == 1 {
			return p.Front().Value.(Primitive)
		} else {
			return NewGrid(p, false)
		}
	}
	return prim
}
