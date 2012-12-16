package render

import (
	//"math"
	"ray/accel"
	"ray/geom"
	//"ray/light"
	"ray/light/spectrum"
	//"ray/shape"
	"ray/world"
)

type Renderer interface {
	Render(world.World)
	Li(wor world.World, ray geom.Ray, computeIsect, computeTransmittance bool) (spectrum.RGBSpectrum, accel.Intersection, spectrum.RGBSpectrum)
	Transmittance(wor world.World, ray geom.Ray) spectrum.RGBSpectrum
}
