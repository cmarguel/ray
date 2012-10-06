package shape

import (
	"container/list"
	"ray/geom"
)

type Shape interface {
	Intersect(ray *geom.Ray) (*DifferentialGeometry, geom.Color, bool)
	WorldBound() geom.BBox
	Refine(*list.List)
}
