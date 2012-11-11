package render

import (
	"math"
	"ray/accel"
	"ray/geom"
	//"ray/light"
	"ray/light/spectrum"
	"ray/shape"
	"ray/world"
)

type Renderer interface {
	Render(world.World)
	Li(wor world.World, ray geom.Ray, computeIsect, computeTransmittance bool) (spectrum.RGBSpectrum, accel.Intersection, spectrum.RGBSpectrum)
	Transmittance(wor world.World, ray geom.Ray) spectrum.RGBSpectrum
}

func evaluateRadiance(wor world.World, dg *shape.DifferentialGeometry) spectrum.RGBSpectrum {
	spec := spectrum.NewRGBSpectrum(0.0)
	for _, light := range wor.Lights {
		spec = spec.Plus(light.SampleL(dg.P))
	}
	return spec.Clamp(0., math.Inf(1))
}
