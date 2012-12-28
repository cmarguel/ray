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

func NewPerspective(c2w, proj mmath.Transform, window []float64, film film.ImageFilm) Perspective {
	proj := NewProjective(c2w, proj, window, film)
	r2c := proj.RasterToCamera
	dx := r2c.Apply(geom.NewVector3(1, 0, 0)).Minus(r2c.Apply(geom.NewVector3(0, 0, 0)))
	dy := r2c.Apply(geom.NewVector3(0, 1, 0)).Minus(r2c.Apply(geom.NewVector3(0, 0, 0)))

	return Perspective{film, dx, dy}
}

func (p Perspective) GenerateRay(sample sampler.Sample) geom.Ray {
	ras := geom.NewVector3(sample.ImageX, sample.ImageY, 0)
	pCamera := p.RasterToCamera.Apply(ras)

	ray := geom.NewRay(geom.NewVector3(0, 0, 0), pCamera)
	// ray.Time = mmath.Lerp(sample.Time, v1, v2)
	ray = p.CameraToWorld.Apply(ray)

	return ray
}

func (p Perspective) GetPos() geom.Vector3 {

}

func (p Perspective) Film() film.Film {

}
