package geom

import (
	"math"
)

type Point Vector3

type Ray struct {
	Origin    Vector3
	Direction Vector3

	MinT *float64
	MaxT *float64
	Time *float64
}

func NewRay(origin, direction Vector3) Ray {
	ray := Ray{origin, direction, new(float64), new(float64), new(float64)}
	*ray.MinT = 0.
	*ray.MaxT = math.Inf(1)
	*ray.Time = 0.

	return ray
}

func (r Ray) At(t float64) Vector3 {
	return r.Origin.Plus(r.Direction.Times(t))
}
