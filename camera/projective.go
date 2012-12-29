package camera

import (
	"ray/camera/film"
	"ray/mmath"
)

type Projective struct {
	CameraToScreen mmath.Transform
	RasterToCamera mmath.Transform
	CameraToWorld  mmath.Transform
	ScreenToRaster mmath.Transform
	RasterToScreen mmath.Transform

	film film.ImageFilm
}

func NewProjective(c2w, proj mmath.Transform, window []float64, film film.ImageFilm) Projective {
	c2s := proj

	s2r := mmath.NewTransform().
		Scale(float64(film.ResolutionX), float64(film.ResolutionY), 1.).
		Scale(1./(window[1]-window[0]), 1./(window[2]-window[3]), 1.).
		Translate(-window[0], -window[3], 0.)

	r2s := s2r.Inverse()

	r2c := c2s.Inverse().Times(r2s)
	return Projective{c2s, r2c, c2w, s2r, r2s, film}
}

func (p Projective) Film() film.Film {
	return p.film
}
