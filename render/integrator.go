package render

import (
	//"ray/light"
	"ray/light/spectrum"
	"ray/world"
)

type Integrator interface {
	RequestSamples(world.World)
}

type SurfaceIntegrator interface {
	Li(world.World) spectrum.RGBSpectrum
}

type DirectLightingIntegrator struct {
}

func (d DirectLightingIntegrator) RequestSamples(wor world.World) {
	//for i, l := wor.Lights {

	//}
}
