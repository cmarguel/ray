package light

import (
	"ray/geom"
	"ray/light/spectrum"
)

type PointLight struct {
	Pos       geom.Vector3
	Intensity spectrum.RGBSpectrum

	nSamples int
}

func NewPointLight(x, y, z, r, g, b float64) PointLight {
	return PointLight{geom.NewVector3(x, y, z), spectrum.FromRGB(r, g, b), 1}
}

func (l PointLight) SampleL(point geom.Vector3) spectrum.RGBSpectrum {
	wi := point.Minus(l.Pos)
	normalizer := 1. / wi.MagnitudeSquared()
	return l.Intensity.TimesC(normalizer)
}

func (l PointLight) IsDeltaLight() bool {
	return true
}

func (l PointLight) NSamples() int {
	return l.nSamples
}
