package shape

import (
	"ray/geom"
)

type Shape interface {
	Intersect(ray geom.Ray) (geom.Vector3, geom.Color, bool)
}
