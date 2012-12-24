package util

import (
	"math/rand"
)

func StratifiedSample1D(samp []float64, nSamples int, jitter bool) {
	inv := 1. / float64(nSamples)
	for i := range samp {
		delta := 0.5
		if jitter {
			delta = rand.Float64()
		}
		samp[i] = (float64(i) + delta) * inv
	}
}

func StratifiedSample2D(samp []float64, nx, ny int, jitter bool) {
	dx, dy := 1./float64(nx), 1./float64(ny)
	i := 0
	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			jx, jy := 0.5, 0.5
			if jitter {
				jx, jy = rand.Float64(), rand.Float64()
			}
			samp[i] = (float64(x) + jx) * dx
			i++
			samp[i] = (float64(y) + jy) * dy
			i++
		}
	}
}
