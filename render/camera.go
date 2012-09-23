package render

import (
	"ray/geom"
)

type Viewport struct {
	Width  float64
	Height float64
	Depth  float64
}

type Camera struct {
	Eye      geom.Vector3
	Viewport Viewport
}

func NewCamera(w, h, d float64) Camera {
	eye := geom.NewVector3(0, 0, 0)
	viewport := Viewport{w, h, d}
	return Camera{eye, viewport}
}
