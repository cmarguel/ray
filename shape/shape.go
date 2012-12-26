package shape

import (
	"container/list"
	"ray/geom"
)

type Shape interface {
	CanIntersect() bool
	Intersect(ray *geom.Ray) (*DifferentialGeometry, float64, float64, bool)
	IntersectP(ray geom.Ray) bool
	WorldBound() geom.BBox
	Refine(*list.List)
}
