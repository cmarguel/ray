package sampler

import (
	"math"
)

type Filter interface {
	XWidth() float64
	YWidth() float64
	InvXWidth() float64
	InvYWidth() float64
	Evaluate(x, y float64) float64
}

type BaseFilter struct {
	xWidth, yWidth, invXWidth, invYWidth float64
}

func (b BaseFilter) XWidth() float64 {
	return b.xWidth
}

func (b BaseFilter) YWidth() float64 {
	return b.yWidth
}

func (b BaseFilter) InvXWidth() float64 {
	return b.invXWidth
}

func (b BaseFilter) InvYWidth() float64 {
	return b.invYWidth
}

type GaussianFilter struct {
	BaseFilter
	alpha, expX, expY float64
}

func NewGaussianFilter(xw, yw, a float64) GaussianFilter {
	expX := math.Exp(-a * xw * xw)
	expY := math.Exp(-a * yw * yw)
	return GaussianFilter{BaseFilter{xw, yw, 1 / xw, 1 / yw}, a, expX, expY}
}

func (g GaussianFilter) gaussian(d, expv float64) float64 {
	return math.Max(0., math.Exp(-g.alpha*d*d)-expv)
}

func (g GaussianFilter) Evaluate(x, y float64) float64 {
	return g.gaussian(x, g.expX) * g.gaussian(y, g.expY)
}
