package light

import (
	"ray/geom"
	"ray/light/spectrum"
)

type Light interface {
	SampleL(point geom.Vector3) spectrum.RGBSpectrum
}
