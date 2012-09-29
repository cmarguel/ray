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

func (sp1 RGBSpectrum) TimesC(t float64) RGBSpectrum {
	sp1.Vals = sp1.copyVals()
	for i := range sp1.Vals {
		sp1.Vals[i] *= t
	}
	return sp1
}

func (sp1 RGBSpectrum) Plus(sp2 RGBSpectrum) RGBSpectrum {
	sp1.Vals = sp1.copyVals()
	for i, v := range sp2.Vals {
		sp1.Vals[i] += v
	}
	return sp1
}

func (sp RGBSpectrum) Clamp(low, high float64) RGBSpectrum {
	sp.Vals = sp.copyVals()
	for i, v := range sp.Vals {
		if v < low {
			sp.Vals[i] = low
		} else if v > high {
			sp.Vals[i] = high
		}
	}
	return sp
}
