package geom

import (
	"math"
)

type Point Vector3
type Color Vector3

type Vertex struct {
	P Vector3
	C Color
}

type Ray struct {
	Origin    Vector3
	Direction Vector3

	MinT *float64
	MaxT *float64
}

func NewRay(origin, direction Vector3) Ray {
	ray := Ray{origin, direction, new(float64), new(float64)}
	*ray.MinT = 0.
	*ray.MaxT = math.Inf(1)

	return ray
}
