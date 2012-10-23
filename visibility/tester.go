package visibility

import (
	"ray/geom"
)

type Tester struct {
	R geom.Ray
}

func NewTester() *Tester {
	return new(Tester)
}

func (t *Tester) SetSegment(p1, p2 geom.Vector3, eps1, eps2, time float64) {
	dist := p1.Minus(p2).Magnitude()
	ray := geom.NewRay(p1, p2)
	*ray.MinT = eps1
	*ray.MaxT = dist * (1. - eps2)
	// ray.Time = time

	t.R = ray
}
