package light

import (
	"ray/geom"
	"ray/light/spectrum"
)

type Light interface {
	NSamples() int
	IsDeltaLight() bool
	SampleL(point geom.Vector3) spectrum.RGBSpectrum
}
