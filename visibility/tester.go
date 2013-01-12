package visibility

import (
	"math"
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
	ray := geom.NewRay(p1, p2.Minus(p1).Times(1/dist))
	*ray.MinT = eps1
	*ray.MaxT = dist * (1. - eps2)
	*ray.Time = time

	t.R = ray
}

func (t *Tester) SetRay(o, d geom.Vector3, eps, time float64) {
	t.R = geom.NewRay(o, d)
	*t.R.MinT = eps
	*t.R.MaxT = math.Inf(1)
	*t.R.Time = time
}
