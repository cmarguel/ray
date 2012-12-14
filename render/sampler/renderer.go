package sampler

import (
	"ray/accel"
	"ray/integrator"
	"ray/light/spectrum"
	"ray/world"
)

type Renderer struct {
	SurfaceIntegrator integrator.SurfaceIntegrator
}

func NewRenderer(surface integrator.SurfaceIntegrator) Renderer {
	return Renderer{surface}
}

func (r Renderer) Li(wor world.World, isect accel.Intersection) spectrum.RGBSpectrum {
	return r.SurfaceIntegrator.Li(wor, isect)
}
