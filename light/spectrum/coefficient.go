package spectrum

type CoefficientSpectrum struct {
	Vals []float64
}

func NewCoefficientSpectrum(nSamples int, v float64) CoefficientSpectrum {
	vals := make([]float64, nSamples)
	for i := range vals {
		vals[i] = v
	}

	return CoefficientSpectrum{vals}
}

func (sp CoefficientSpectrum) copyVals() []float64 {
	vals := make([]float64, len(sp.Vals))
	copy(vals, sp.Vals)
	return vals
}

func (sp1 CoefficientSpectrum) Plus(sp2 CoefficientSpectrum) CoefficientSpectrum {
	sp1.Vals = sp1.copyVals()
	for i, v := range sp2.Vals {
		sp1.Vals[i] += v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Minus(sp2 CoefficientSpectrum) CoefficientSpectrum {
	sp1.Vals = sp1.copyVals()
	for i, v := range sp2.Vals {
		sp1.Vals[i] -= v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Times(sp2 CoefficientSpectrum) CoefficientSpectrum {
	sp1.Vals = sp1.copyVals()
	for i, v := range sp2.Vals {
		sp1.Vals[i] *= v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) TimesC(t float64) CoefficientSpectrum {
	sp1.Vals = sp1.copyVals()
	for i := range sp1.Vals {
		sp1.Vals[i] *= t
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Neg() CoefficientSpectrum {
	sp1.Vals = sp1.copyVals()
	for i, v := range sp1.Vals {
		sp1.Vals[i] = -v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Lerp(t float64, sp2 CoefficientSpectrum) CoefficientSpectrum {
	return sp1.TimesC(1. - t).Plus(sp2.TimesC(t))
}

func (sp CoefficientSpectrum) Clamp(low, high float64) CoefficientSpectrum {
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
