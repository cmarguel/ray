package sampler

import (
	"ray/camera"
	"ray/camera/film"
	"ray/geom"
	"ray/integrator"
	"ray/light/spectrum"
	"ray/world"
)

type Renderer struct {
	SurfaceIntegrator integrator.SurfaceIntegrator

	camera camera.Camera
}

func NewRenderer(cam camera.Camera, fil film.Film, surface integrator.SurfaceIntegrator) Renderer {
	return Renderer{surface, cam}
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
