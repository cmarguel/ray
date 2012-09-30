package accel

import (
	"math"
	"ray/geom"
)

type BBox struct {
	Min geom.Vector3
	Max geom.Vector3
}

func NewBBoxEmpty() BBox {
	plus := math.Inf(1)
	minus := math.Inf(-1)

	max := geom.NewVector3(minus, minus, minus)
	min := geom.NewVector3(plus, plus, plus)
	return BBox{min, max}
}

func NewBBox(p1, p2 geom.Vector3) BBox {
	min := geom.NewVector3(math.Min(p1.X, p2.X), math.Min(p1.Y, p2.Y), math.Min(p1.Z, p2.Z))
	max := geom.NewVector3(math.Max(p1.X, p2.X), math.Max(p1.Y, p2.Y), math.Max(p1.Z, p2.Z))
	return BBox{min, max}
}

func (b BBox) AddPoint(p geom.Vector3) BBox {
	min := geom.NewVector3(math.Min(b.Min.X, p.X), math.Min(b.Min.Y, p.Y), math.Min(b.Min.Z, p.Z))
	max := geom.NewVector3(math.Max(b.Max.X, p.X), math.Max(b.Max.Y, p.Y), math.Max(b.Max.Z, p.Z))
	return BBox{min, max}
}

func (b BBox) Union(d BBox) BBox {
	min := geom.NewVector3(math.Min(b.Min.X, d.Min.X), math.Min(b.Min.Y, d.Min.Y), math.Min(b.Min.Z, d.Min.Z))
	max := geom.NewVector3(math.Max(b.Max.X, d.Max.X), math.Max(b.Max.Y, d.Max.Y), math.Max(b.Max.Z, d.Max.Z))
	return BBox{min, max}
}

func (b BBox) Expand(delta float64) BBox {
	d := geom.NewVector3(delta, delta, delta)
	b.Min = b.Min.Minus(d)
	b.Max = b.Max.Plus(d)
	return b
}

func (b BBox) Inside(p geom.Vector3) bool {
	return p.X >= b.Min.X && p.X <= b.Max.X &&
		p.Y >= b.Min.Y && p.Y <= b.Max.Y &&
		p.Z >= b.Min.Z && p.Z <= b.Max.Z

}

func getNewInterval(t0, t1, tNear, tFar float64) (float64, float64, bool) {
	if tNear > tFar {
		tNear, tFar = tFar, tNear
	}
	if tNear > t0 {
		t0 = tNear
	}
	if tFar < t1 {
		t1 = tFar
	}

	if t0 > t1 {
		return t0, t1, false
	}

	return t0, t1, true
}

func (b BBox) IntersectP(ray geom.Ray) (float64, float64, bool) {
	t0 := *ray.MinT
	t1 := *ray.MaxT
	status := false

	invRayDirX := 1. / ray.Direction.X
	tNearX := (b.Min.X - ray.Origin.X) * invRayDirX
	tFarX := (b.Max.X - ray.Origin.X) * invRayDirX
	t0, t1, status = getNewInterval(t0, t1, tNearX, tFarX)
	if !status {
		return t0, t1, false
	}

	invRayDirY := 1. / ray.Direction.Y
	tNearY := (b.Min.Y - ray.Origin.Y) * invRayDirY
	tFarY := (b.Max.Y - ray.Origin.Y) * invRayDirY
	t0, t1, status = getNewInterval(t0, t1, tNearY, tFarY)
	if !status {
		return t0, t1, false
	}

	invRayDirZ := 1. / ray.Direction.Z
	tNearZ := (b.Min.Z - ray.Origin.Z) * invRayDirZ
	tFarZ := (b.Max.Z - ray.Origin.Z) * invRayDirZ
	t0, t1, status = getNewInterval(t0, t1, tNearZ, tFarZ)
	if !status {
		return t0, t1, false
	}

	return t0, t1, true
}
