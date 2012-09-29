package spectrum

type RGBSpectrum struct {
	CoefficientSpectrum
}

func NewRGBSpectrum(v float64) RGBSpectrum {
	return RGBSpectrum{NewCoefficientSpectrum(3, v)}
}

func FromRGB(r, g, b float64) RGBSpectrum {
	sp := NewCoefficientSpectrum(3, 0)
	sp.Vals[0] = r
	sp.Vals[1] = g
	sp.Vals[2] = b
	return RGBSpectrum{sp}
}

func (rgb RGBSpectrum) ToRGB() (r, g, b float64) {
	r = rgb.Vals[0]
	g = rgb.Vals[1]
	b = rgb.Vals[2]
	return r, g, b
}
