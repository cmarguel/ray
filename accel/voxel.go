package accel

import (
	"container/list"
	"ray/geom"
	"sync"
)

type Voxel struct {
	Primitives      *list.List
	AllCanIntersect bool
	mutex           *sync.Mutex
}

func NewVoxel(p Primitive) Voxel {
	primitives := list.New()
	primitives.PushBack(p)
	return Voxel{primitives, false, new(sync.Mutex)}
}

func (v Voxel) AddPrimitive(p Primitive) {
	v.Primitives.PushBack(p)
}

func (v Voxel) Intersect(ray *geom.Ray) (Intersection, bool) {
	if !v.AllCanIntersect {
		v.mutex.Lock()
		for e := v.Primitives.Front(); e != nil; e = e.Next() {
			prim := e.Value.(Primitive)
			refined := v.refine(prim)
			e.Value = refined
		}
		v.AllCanIntersect = true
		v.mutex.Unlock()
	}

	return v.findIntersections(ray)
}

func (v Voxel) IntersectP(ray geom.Ray) bool {
	if !v.AllCanIntersect {
		v.mutex.Lock()
		for e := v.Primitives.Front(); e != nil; e = e.Next() {
			prim := e.Value.(Primitive)
			refined := v.refine(prim)
			e.Value = refined
		}
		v.mutex.Unlock()
		v.AllCanIntersect = true
	}

	return v.findIntersectionsP(ray)
}

func (v Voxel) findIntersectionsP(ray geom.Ray) bool {
	for e := v.Primitives.Front(); e != nil; e = e.Next() {
		prim := e.Value.(Primitive)
		found := prim.IntersectP(ray)
		if found {
			return true
		}
	}
	return false
}

func (v Voxel) findIntersections(ray *geom.Ray) (Intersection, bool) {
	hitSomething := false
	var intersection = *new(Intersection)

	for e := v.Primitives.Front(); e != nil; e = e.Next() {
		prim := e.Value.(Primitive)
		isect, found := prim.Intersect(ray)
		if found {
			hitSomething = true
			intersection = isect
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
