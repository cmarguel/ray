package film

import (
	"ray/light/spectrum"
	"ray/render/sampler"
)

type Film interface {
	W() int
	H() int

	AddSample(sampler.Sample, spectrum.RGBSpectrum)
}

type BaseFilm struct {
	ResolutionX int
	ResolutionY int
}

func (b BaseFilm) W() int {
	return b.ResolutionX
}

func (b BaseFilm) H() int {
	return b.ResolutionY
}

type ImageFilm struct {
	BaseFilm
	Pixels [][]*Pixel

	writer chan PixPart
}

func NewImageFilm(w, h int) ImageFilm {
	p := make([][]*Pixel, w)
	for x := 0; x < w; x++ {
		p[x] = make([]*Pixel, h)
		for y := 0; y < h; y++ {
			pix := NewPixel()
			p[x][y] = &pix
		}
	}

	img := ImageFilm{BaseFilm{w, h}, p, make(chan PixPart)}
	go img.acceptWrites()

	return img
}

type PixPart struct {
	x, y   int
	lxyz   []float64
	weight float64
}

func (img ImageFilm) Add(x, y int, lxyz []float64, weight float64) {
	if x >= 0 && y >= 0 && x < img.ResolutionX && y < img.ResolutionY {
		img.writer <- PixPart{x, y, lxyz, weight}
	}
}

// Use channels to write to the underlying image so that multiple threads can add samples concurrently.
func (img ImageFilm) acceptWrites() {
	for p := range img.writer {
		x, y := p.x, p.y
		img.Pixels[x][y].Lxyz[0] += p.weight * p.lxyz[0]
		img.Pixels[x][y].Lxyz[1] += p.weight * p.lxyz[1]
		img.Pixels[x][y].Lxyz[2] += p.weight * p.lxyz[2]
		img.Pixels[x][y].WeightSum += p.weight
	}
}

func (img ImageFilm) AddSample(sample sampler.Sample, rgb spectrum.RGBSpectrum) {
	x := int(sample.ImageX)
	y := int(sample.ImageY)
	img.Add(x, y, rgb.Vals, 1.)
}
