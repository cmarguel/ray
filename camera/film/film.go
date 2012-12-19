package film

import (
	"math"
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

	filter sampler.Filter
	writer chan PixPart
}

func NewImageFilm(w, h int, filt sampler.Filter) ImageFilm {
	p := make([][]*Pixel, w)
	for x := 0; x < w; x++ {
		p[x] = make([]*Pixel, h)
		for y := 0; y < h; y++ {
			pix := NewPixel()
			p[x][y] = &pix
		}
	}

	img := ImageFilm{BaseFilm{w, h}, p, filt, make(chan PixPart)}
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
	dx := sample.ImageX - 0.5
	dy := sample.ImageY - 0.5
	x0 := int(math.Ceil(dx - img.filter.XWidth()))
	x1 := int(math.Ceil(dx + img.filter.XWidth()))
	y0 := int(math.Ceil(dy - img.filter.YWidth()))
	y1 := int(math.Ceil(dy + img.filter.YWidth()))

	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			weight := img.filter.Evaluate(float64(x)-sample.ImageX, float64(y)-sample.ImageY)
			img.Add(x, y, rgb.Vals, weight)
		}
	}

}
