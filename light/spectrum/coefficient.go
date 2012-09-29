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

func (sp1 CoefficientSpectrum) Plus(sp2 CoefficientSpectrum) CoefficientSpectrum {
	for i, v := range sp2.Vals {
		sp1.Vals[i] += v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Minus(sp2 CoefficientSpectrum) CoefficientSpectrum {
	for i, v := range sp2.Vals {
		sp1.Vals[i] -= v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Times(sp2 CoefficientSpectrum) CoefficientSpectrum {
	for i, v := range sp2.Vals {
		sp1.Vals[i] *= v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) TimesC(t float64) CoefficientSpectrum {
	for i := range sp1.Vals {
		sp1.Vals[i] *= t
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Neg() CoefficientSpectrum {
	for i, v := range sp1.Vals {
		sp1.Vals[i] = -v
	}
	return sp1
}

func (sp1 CoefficientSpectrum) Lerp(t float64, sp2 CoefficientSpectrum) CoefficientSpectrum {
	return sp1.TimesC(1. - t).Plus(sp2.TimesC(t))
}
