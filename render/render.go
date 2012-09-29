package render

import (
	"math"
	"ray/geom"
	//"ray/light"
	"ray/light/spectrum"
	"ray/world"
)

func evaluateRadiance(wor world.World, p geom.Vector3) spectrum.RGBSpectrum {
	spec := spectrum.NewRGBSpectrum(0.0)
	for _, light := range wor.Lights {
		spec = spec.Plus(light.SampleL(p))
	}
	return spec.Clamp(0., math.Inf(1))
}
