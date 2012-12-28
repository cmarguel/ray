package integrator

import (
	"math"
	"ray/accel"
	"ray/light/spectrum"
	"ray/world"
)

type WhittedIntegrator struct {
	MaxDepth int
}

func NewWhitted() WhittedIntegrator {
	return WhittedIntegrator{5}
}

func (w WhittedIntegrator) Li(wor world.World, isect accel.Intersection) spectrum.RGBSpectrum {
	spec := spectrum.NewRGBSpectrum(0.1)

	for _, light := range wor.Lights {
		spectrum, _, tester := light.SampleL(isect.DiffGeom.P, isect.RayEpsilon, 0)
		if (!spectrum.IsBlack()) && wor.Unoccluded(*tester) {
			spec = spec.Plus(spectrum)
		}
	}
	return spec.Clamp(0., math.Inf(1))
}

func (w WhittedIntegrator) RequestSamples(world.World) {

}
