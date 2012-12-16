package film

type Film interface {
}

type BaseFilm struct {
	ResolutionX int
	ResolutionY int
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
		img.Pixels[x][y].Lxyz[0] = p.lxyz[0]
		img.Pixels[x][y].Lxyz[1] = p.lxyz[1]
		img.Pixels[x][y].Lxyz[2] = p.lxyz[2]
		img.Pixels[x][y].WeightSum += p.weight
	}
}
