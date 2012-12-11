package light

import (
	"ray/geom"
	"ray/light/spectrum"
	"ray/visibility"
)

type Light interface {
	NSamples() int
	IsDeltaLight() bool
	SampleL(geom.Vector3, float64, float64) (spectrum.RGBSpectrum, geom.Vector3, *visibility.Tester)
}
