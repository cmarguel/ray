package light

import (
	"ray/geom"
	"ray/light/spectrum"
	"ray/visibility"
)

type PointLight struct {
	Pos       geom.Vector3
	Intensity spectrum.RGBSpectrum

	nSamples int
}

func NewPointLight(x, y, z, r, g, b float64) PointLight {
	return PointLight{geom.NewVector3(x, y, z), spectrum.FromRGB(r, g, b), 1}
}

func (l PointLight) SampleL(point geom.Vector3, pEpsilon, time float64) (spectrum.RGBSpectrum, geom.Vector3, *visibility.Tester) {
	wi := point.Minus(l.Pos)
	normalizer := 1. / wi.MagnitudeSquared()
	tester := visibility.NewTester()
	tester.SetSegment(point, l.Pos, pEpsilon, 0., time)
	return l.Intensity.TimesC(normalizer), l.Pos.Minus(point).Normalized(), tester
}

func (l PointLight) IsDeltaLight() bool {
	return true
}

func (l PointLight) NSamples() int {
	return l.nSamples
}
