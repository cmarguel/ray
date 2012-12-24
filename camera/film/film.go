package film

import (
	"math"
	"ray/light/spectrum"
	"ray/render/sampler"
)

const size = 16

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

	filter      sampler.Filter
	writer      chan PixPart
	filterTable []float64
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

	table := precomputeFilter(size, filt)

	img := ImageFilm{BaseFilm{w, h}, p, filt, make(chan PixPart, w*h), table}
	go img.acceptWrites()

	return img
}

func precomputeFilter(size int, filt sampler.Filter) []float64 {
	table := make([]float64, size*size)
	i := 0
	for y := 0; i < size; y++ {
		fy := (float64(y) + 0.5) * (filt.YWidth() / float64(size))
		for x := 0; x < size; x++ {
			fx := (float64(x) + 0.5) * (filt.XWidth() / float64(size))
			table[i] = filt.Evaluate(fx, fy)
			i++
		}
	}
	return table
}

type PixPart struct {
	x, y   int
	lxyz   []float64
	weight float64
}

func (img ImageFilm) AddConcurrent(x, y int, lxyz []float64, weight float64) {
	if x >= 0 && y >= 0 && x < img.ResolutionX && y < img.ResolutionY {
		img.writer <- PixPart{x, y, lxyz, weight}
	}
}

// Use channels to write to the underlying image so that multiple threads can add samples concurrently.
func (img ImageFilm) acceptWrites() {
	for p := range img.writer {
		img.addPixel(p.x, p.y, p.lxyz[0], p.lxyz[1], p.lxyz[2], p.weight)
	}
}

func (img ImageFilm) addPixel(x, y int, lx, ly, lz, weight float64) {
	if x >= 0 && y >= 0 && x < img.ResolutionX && y < img.ResolutionY {
		img.Pixels[x][y].Lxyz[0] += weight * lx
		img.Pixels[x][y].Lxyz[1] += weight * ly
		img.Pixels[x][y].Lxyz[2] += weight * lz
		img.Pixels[x][y].WeightSum += weight
	}
}

func (img ImageFilm) computeTableOffsets(dx, dy float64, x0, x1, y0, y1 int) ([]int, []int) {
	ifx := make([]int, x1-x0+1)
	for x := x0; x <= x1; x++ {
		fx := math.Abs((float64(x) - dx) * img.filter.InvXWidth() * size)
		ifx[x-x0] = int(math.Min(math.Floor(fx), size-1))
	}
	ify := make([]int, y1-y0+1)
	for y := y0; y <= y1; y++ {
		fy := math.Abs((float64(y) - dy) * img.filter.InvYWidth() * size)
		ify[y-y0] = int(math.Min(math.Floor(fy), size-1))
	}
	return ifx, ify
}

func (img ImageFilm) AddSample(sample sampler.Sample, rgb spectrum.RGBSpectrum) {
	dx := sample.ImageX - 0.5
	dy := sample.ImageY - 0.5
	x0 := int(math.Ceil(dx - img.filter.XWidth()))
	x1 := int(math.Ceil(dx + img.filter.XWidth()))
	y0 := int(math.Ceil(dy - img.filter.YWidth()))
	y1 := int(math.Ceil(dy + img.filter.YWidth()))

	//ifx, ify := img.computeTableOffsets(dx, dy, x0, x1, y0, y1)
	//sync := img.filter.XWidth() > 0.5 || img.filter.YWidth() > 0.5

	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			//offset := ify[y-y0]*size + ifx[x-x0]
			//weight := img.filterTable[offset]
			weight := img.filter.Evaluate(float64(x)-sample.ImageX, float64(y)-sample.ImageY)
			//if sync {
			img.AddConcurrent(x, y, rgb.Vals, weight)
			//} else {
			//	img.addPixel(x, y, rgb.Vals[0], rgb.Vals[1], rgb.Vals[2], weight)
			//}
		}
	}

}
