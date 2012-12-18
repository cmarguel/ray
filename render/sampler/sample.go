package sampler

type CameraSample struct {
	ImageX, ImageY, LensU, LensV, Time float64
}

type Sample struct {
	CameraSample
}

func NewSample(x, y, u, v, t float64) Sample {
	return Sample{CameraSample{x, y, u, v, t}}
}
