package integrator

import (
	//"ray/light"
	"ray/accel"
	"ray/light/spectrum"
	"ray/world"
)

type Integrator interface {
	RequestSamples(world.World)
}

type SurfaceIntegrator interface {
	Integrator
	Li(world.World, accel.Intersection) spectrum.RGBSpectrum
}

type DirectLightingIntegrator struct {
}

func (d DirectLightingIntegrator) RequestSamples(wor world.World) {
	//for i, l := wor.Lights {

	//}
}
