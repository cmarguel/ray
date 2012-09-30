package accel

import (
	"container/list"
	"ray/geom"
)

type Aggregate struct {
}

func (a Aggregate) WorldBound() geom.BBox {
	return *new(geom.BBox)
}

func (a Aggregate) CanIntersect() bool {
	return false
}

func (a Aggregate) Intersect(geom.Ray) (Intersection, bool) {
	return *new(Intersection), false
}

func (a Aggregate) Refine(*list.List) {

}
