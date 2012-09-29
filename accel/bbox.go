package accel

import (
	"math"
	"ray/geom"
)

type BBox struct {
	Min geom.Vector3
	Max geom.Vector3
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

func (b BBox) IntersectP(ray geom.Ray) (float64, float64, bool) {
	return 0, 0, false
}
