package sampler

import (
	"ray/mmath"
	"ray/util"
)

type Stratified struct {
	BaseSampler

	xPos, yPos         int
	xSamples, ySamples int
	jitter             bool
	buffer             []float64
}

func NewStratified(xStart, xEnd, yStart, yEnd, xs, ys int, jitter bool, sOpen, sClose float64) *Stratified {
	base := BaseSampler{xStart, xEnd, yStart, yEnd, xs * ys, sOpen, sClose}
	sampler := &Stratified{base, xStart, yStart, xs, ys, jitter, make([]float64, 5*xs*ys)}

	return sampler
}

func (s *Stratified) GetMoreSamples() ([]Sample, bool) {
	if s.yPos == s.YEnd {
		return []Sample{}, false
	}
	nSamples := s.xSamples * s.ySamples

	samples := s.generateSamples(nSamples)

	s.nextPixel()

	return samples, true
}

func (s *Stratified) generateSamples(nSamples int) []Sample {
	imageSamples := s.buffer[0 : 2*nSamples]
	lensSamples := s.buffer[2*nSamples : 4*nSamples]
	timeSamples := s.buffer[4*nSamples:]

	util.StratifiedSample2D(imageSamples, s.xSamples, s.ySamples, s.jitter)
	util.StratifiedSample2D(lensSamples, s.xSamples, s.ySamples, s.jitter)
	util.StratifiedSample1D(timeSamples, nSamples, s.jitter)

	s.adjustImageSamples(imageSamples)
	util.ShuffleFloat64(lensSamples)
	util.ShuffleFloat64(timeSamples)

	samples := make([]Sample, nSamples)
	for i := range samples {
		samples[i].ImageX = imageSamples[2*i]
		samples[i].ImageY = imageSamples[2*i+1]
		samples[i].LensU = lensSamples[2*i]
		samples[i].LensV = lensSamples[2*i+1]
		samples[i].Time = mmath.Lerp(timeSamples[i], s.ShutterOpen, s.ShutterClose)
	}
	return samples
}

func (s *Stratified) adjustImageSamples(samples []float64) {
	for i := 0; i < len(samples); i += 2 {
		samples[i] += float64(s.xPos)
		samples[i+1] += float64(s.yPos)
	}
}

func (s *Stratified) nextPixel() {
	s.xPos += 1
	if s.xPos == s.XEnd {
		s.xPos = s.XStart
		s.yPos += 1
	}
}

func (s *Stratified) GetSubSampler(num, count int) Sampler {
	x0, x1, y0, y1 := s.computeSubWindow(num, count)
	if x0 == x1 || y0 == y1 {
		return nil
	}
	return NewStratified(x0, x1, y0, y1, s.xSamples, s.ySamples, s.jitter, s.ShutterOpen, s.ShutterClose)
}
