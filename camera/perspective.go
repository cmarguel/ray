package camera

import (
	"ray/camera/film"
	"ray/geom"
	"ray/mmath"
	"ray/render/sampler"
)

type Perspective struct {
	Projective
	DxCamera, DyCamera geom.Vector3
}

func NewPerspective(c2w mmath.Transform, window []float64, fov float64, film film.ImageFilm) Perspective {
	persp := mmath.NewTransform().Perspective(fov, 1e-2, 1000.)
	proj := NewProjective(c2w, persp, window, film)
	r2c := proj.RasterToCamera
	dx := r2c.Apply(geom.NewVector3(1, 0, 0)).Minus(r2c.Apply(geom.NewVector3(0, 0, 0)))
	dy := r2c.Apply(geom.NewVector3(0, 1, 0)).Minus(r2c.Apply(geom.NewVector3(0, 0, 0)))

	return Perspective{proj, dx, dy}
}

func (p Perspective) GenerateRay(sample sampler.Sample) geom.Ray {
	ras := geom.NewVector3(sample.ImageX, sample.ImageY, 0)
	pCamera := p.RasterToCamera.Apply(ras)

	ray := geom.NewRay(geom.NewVector3(0, 0, 0), pCamera)
	// ray.Time = mmath.Lerp(sample.Time, v1, v2)
	ray = p.CameraToWorld.ApplyToRay(ray)

	return ray
}

func (p Perspective) Film() film.Film {
	return p.film
}
