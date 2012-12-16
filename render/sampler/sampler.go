package sampler

import (
	"math"
	"ray/mmath"
)

type Sampler interface {
	GetMoreSamples() ([]Sample, bool)
	GetSubSampler(num, count int) Sampler
}

type BaseSampler struct {
	XStart, XEnd              int
	YStart, YEnd              int
	SamplesPerPixel           int
	ShutterOpen, ShutterClose float64
}

func (b BaseSampler) computeSubWindow(num, count int) (int, int, int, int) {
	dx := b.XEnd - b.XStart
	dy := b.YEnd - b.YStart
	nx := count
	ny := 1
	// approximately square; doesn't have to be exact
	for (nx&0x1) == 0 && (2*dx*ny) < (dy*nx) {
		nx >>= 1
		ny <<= 1
	}

	xo := num % nx
	yo := num / nx

	// number of tiles in each direction
	tx0 := float64(xo) / float64(nx)
	tx1 := float64(xo+1) / float64(nx)
	ty0 := float64(yo) / float64(ny)
	ty1 := float64(yo+1) / float64(ny)

	xs := int(math.Floor(mmath.Lerp(tx0, float64(b.XStart), float64(b.XEnd))))
	xe := int(math.Floor(mmath.Lerp(tx1, float64(b.XStart), float64(b.XEnd))))
	ys := int(math.Floor(mmath.Lerp(ty0, float64(b.YStart), float64(b.YEnd))))
	ye := int(math.Floor(mmath.Lerp(ty1, float64(b.YStart), float64(b.YEnd))))

	return xs, xe, ys, ye
}
