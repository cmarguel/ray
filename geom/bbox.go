package geom

import (
	"math"
)

type BBox struct {
	Min Vector3
	Max Vector3
}

func NewBBoxEmpty() BBox {
	plus := math.Inf(1)
	minus := math.Inf(-1)

	max := NewVector3(minus, minus, minus)
	min := NewVector3(plus, plus, plus)
	return BBox{min, max}
}

func NewBBox(points ...Vector3) BBox {
	box := NewBBoxEmpty()
	for _, p := range points {
		box = box.AddPoint(p)
	}
	return box
}

func (b BBox) AddPoint(p Vector3) BBox {
	min := NewVector3(math.Min(b.Min.X, p.X), math.Min(b.Min.Y, p.Y), math.Min(b.Min.Z, p.Z))
	max := NewVector3(math.Max(b.Max.X, p.X), math.Max(b.Max.Y, p.Y), math.Max(b.Max.Z, p.Z))
	return BBox{min, max}
}

func (b BBox) Union(d BBox) BBox {
	min := NewVector3(math.Min(b.Min.X, d.Min.X), math.Min(b.Min.Y, d.Min.Y), math.Min(b.Min.Z, d.Min.Z))
	max := NewVector3(math.Max(b.Max.X, d.Max.X), math.Max(b.Max.Y, d.Max.Y), math.Max(b.Max.Z, d.Max.Z))
	return BBox{min, max}
}

func (b BBox) Expand(delta float64) BBox {
	d := NewVector3(delta, delta, delta)
	b.Min = b.Min.Minus(d)
	b.Max = b.Max.Plus(d)
	return b
}

func (b BBox) Inside(p Vector3) bool {
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

func (b BBox) IntersectP(ray Ray) (float64, float64, bool) {
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

func (b BBox) MaximumExtent() int {
	diag := b.Max.Minus(b.Min)
	if diag.X > diag.Y && diag.X > diag.Z {
		return 0
	} else if diag.Y > diag.Z {
		return 1
	}

	return 2
}
