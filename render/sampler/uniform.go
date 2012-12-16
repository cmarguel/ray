package sampler

type UniformSampler struct {
	BaseSampler

	x, y int
}

func NewUniformSampler(xStart, xEnd, yStart, yEnd, perPix int, open, cl float64) UniformSampler {
	return UniformSampler{BaseSampler{xStart, xEnd, yStart, yEnd, perPix, open, cl}, xStart, yStart}
}

func (s *UniformSampler) GetMoreSamples() ([]Sample, bool) {
	if s.y > s.YEnd {
		return []Sample{}, false
	}
	samp := Sample{CameraSample{float64(s.x), float64(s.y), 0, 0, 0}}
	s.x += 1
	if s.x >= s.XEnd {
		s.x = s.XStart
		s.y += 1
	}
	return []Sample{samp}, true
}

func (s UniformSampler) GetSubSampler(num, count int) Sampler {
	x0, x1, y0, y1 := s.computeSubWindow(num, count)
	if x0 == x1 || y0 == y1 {
		panic("Bad data passed to subsampler.")
	}
	sub := NewUniformSampler(x0, x1, y0, y1, s.SamplesPerPixel, s.ShutterOpen, s.ShutterClose)
	return &sub
}
