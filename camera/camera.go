package camera

import (
	"ray/geom"
	"ray/mmath"
)

type Viewport struct {
	Width  float64
	Height float64
	Depth  float64
}

type CameraSample struct {
	ImageX int
	ImageY int
}

func NewCameraSample(x, y int) CameraSample {
	return CameraSample{x, y}
}

type Camera interface {
	GenerateRay(sample CameraSample) geom.Ray
	GetPos() geom.Vector3
}

type Film struct {
	ResolutionX int
	ResolutionY int
}

type PinholeCamera struct {
	Pos    geom.Vector3
	LookAt geom.Vector3
	Up     geom.Vector3

	Film Film

	camToWorld mmath.Transform
}

func NewPinholeCamera(film Film) PinholeCamera {
	pos := geom.NewVector3(1., 0., 0)
	look := geom.NewVector3(0., 0., 1.)
	up := geom.NewVector3(0, 1., 0)

	tr := mmath.NewTransform().LookAt(pos, look, up)

	return PinholeCamera{pos, look, up, film, tr}
}

func (c PinholeCamera) GenerateRay(sample CameraSample) geom.Ray {
	film := c.Film
	x := float64(sample.ImageX - film.ResolutionX/2)
	y := float64(sample.ImageY - film.ResolutionY/2)
	dir := geom.NewVector3(x, y, 100.)
	transformed := c.camToWorld.Apply(dir)

	origin := c.Pos
	newDir := transformed.Minus(c.Pos)
	normalizer := 1. / newDir.Magnitude()
	newDir = newDir.Times(normalizer).Plus(c.Pos)
	return geom.Ray{origin, newDir}
}

func (p PinholeCamera) GetPos() geom.Vector3 {
	return p.Pos
}
