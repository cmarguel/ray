package camera

import (
	"ray/camera/film"
	"ray/geom"
	"ray/mmath"
	"ray/render/sampler"
)

type Viewport struct {
	Width  float64
	Height float64
	Depth  float64
}

type Camera interface {
	GenerateRay(sample sampler.Sample) geom.Ray
	GetPos() geom.Vector3
	Film() film.Film
}

type PinholeCamera struct {
	Pos    geom.Vector3
	LookAt geom.Vector3
	Up     geom.Vector3

	film film.ImageFilm

	camToWorld mmath.Transform
}

func (c PinholeCamera) Film() film.Film {
	return c.film
}

func NewPinholeCamera(film film.ImageFilm) PinholeCamera {
	pos := geom.NewVector3(0., 0., 0)
	look := geom.NewVector3(0., 0., 1.)
	up := geom.NewVector3(0, 1., 0)

	tr := mmath.NewTransform().LookAt(pos, look, up)

	return PinholeCamera{pos, look, up, film, tr}
}

func (c PinholeCamera) GenerateRay(sample sampler.Sample) geom.Ray {
	film := c.film
	x := sample.ImageX - float64(film.ResolutionX/2)
	y := sample.ImageY - float64(film.ResolutionY/2)
	// Scale both components by just one component of the resolution. I suspect 
	// it would look stretched out if we scaled x by x and y by y.
	dir := geom.NewVector3(x/float64(film.ResolutionY), y/float64(film.ResolutionY), 0.1)
	transformed := c.camToWorld.Apply(dir)

	origin := c.Pos
	newDir := transformed.Minus(c.Pos)
	normalizer := 1. / newDir.Magnitude()
	newDir = newDir.Times(normalizer)
	return geom.NewRay(origin, newDir)
}

func (p PinholeCamera) GetPos() geom.Vector3 {
	return p.Pos
}
