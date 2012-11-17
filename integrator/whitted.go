package integrator

import (
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
	l := spectrum.NewRGBSpectrum(0)

	l = l.Plus(isect.Le())

	return l
}

func (w WhittedIntegrator) RequestSamples(world.World) {

}
