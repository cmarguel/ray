package film

type Pixel struct {
	Lxyz      []float64
	Splat     []float64
	WeightSum float64
}

func NewPixel() Pixel {
	return Pixel{[]float64{0, 0, 0}, []float64{0, 0, 0}, 0}
}
