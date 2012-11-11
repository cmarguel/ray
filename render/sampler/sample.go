package sampler

type CameraSample struct {
	ImageX, ImageY, LensU, LensV, Time float64
}

type Sample struct {
	CameraSample
}
