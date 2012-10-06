package render

import (
	"math"
	//"ray/geom"
	//"ray/light"
	"ray/light/spectrum"
	"ray/shape"
	"ray/world"
)

func evaluateRadiance(wor world.World, dg *shape.DifferentialGeometry) spectrum.RGBSpectrum {
	spec := spectrum.NewRGBSpectrum(0.0)
	for _, light := range wor.Lights {
		spec = spec.Plus(light.SampleL(dg.P))
	}
	return spec.Clamp(0., math.Inf(1))
}
