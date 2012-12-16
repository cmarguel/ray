package sampler

import (
	"ray/geom"
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

func (r Renderer) Li(ray geom.Ray, wor world.World) spectrum.RGBSpectrum {
	isect, found := wor.Aggregate.Intersect(&ray)

	li := spectrum.NewRGBSpectrum(0.)
	if found {
		li = r.SurfaceIntegrator.Li(wor, isect)
	} else { // figure out the color when nothing hit
		for _, l := range wor.Lights {
			li = li.Plus(l.Le(ray))
		}
	}
	return li
}
