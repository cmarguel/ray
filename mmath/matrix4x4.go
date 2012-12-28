package mmath

import (
	"math"
)

type Matrix4x4 struct {
	M [][]float64
}

func NewMatrix4x4(
	t00, t01, t02, t03,
	t10, t11, t12, t13,
	t20, t21, t22, t23,
	t30, t31, t32, t33 float64) Matrix4x4 {
	return Matrix4x4{
		[][]float64{
			[]float64{t00, t01, t02, t03},
			[]float64{t10, t11, t12, t13},
			[]float64{t20, t21, t22, t23},
			[]float64{t30, t31, t32, t33},
		},
	}
}

func (mat Matrix4x4) Transpose() Matrix4x4 {
	f := make([][]float64, 4)
	for i := range mat.M {
		f[i] = make([]float64, 4)
	}
	for i, row := range mat.M {
		for j, v := range row {
			f[j][i] = v
		}
	}
	return Matrix4x4{f}
}

func (mat Matrix4x4) Row(i int) []float64 {
	m := mat.M
	return []float64{m[i][0], m[i][1], m[i][2], m[i][3]}
}

func (mat Matrix4x4) Col(i int) []float64 {
	m := mat.M
	return []float64{m[0][i], m[1][i], m[2][i], m[3][i]}
}

func dotM(v1, v2 Matrix4x4, col, row int) float64 {
	a := v1.Row(col)
	b := v2.Col(row)
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

func (m1 Matrix4x4) Times(m2 Matrix4x4) Matrix4x4 {
	t00 := dotM(m1, m2, 0, 0)
	t01 := dotM(m1, m2, 0, 1)
	t02 := dotM(m1, m2, 0, 2)
	t03 := dotM(m1, m2, 0, 3)

	t10 := dotM(m1, m2, 1, 0)
	t11 := dotM(m1, m2, 1, 1)
	t12 := dotM(m1, m2, 1, 2)
	t13 := dotM(m1, m2, 1, 3)

	t20 := dotM(m1, m2, 2, 0)
	t21 := dotM(m1, m2, 2, 1)
	t22 := dotM(m1, m2, 2, 2)
	t23 := dotM(m1, m2, 2, 3)

	t30 := dotM(m1, m2, 3, 0)
	t31 := dotM(m1, m2, 3, 1)
	t32 := dotM(m1, m2, 3, 2)
	t33 := dotM(m1, m2, 3, 3)

	return Matrix4x4{
		[][]float64{
			[]float64{t00, t01, t02, t03},
			[]float64{t10, t11, t12, t13},
			[]float64{t20, t21, t22, t23},
			[]float64{t30, t31, t32, t33},
		},
	}
}

// Copied wholesale from pbrt. Practically a cut-paste job.
// Sorry; I'm not interested in trying to implement Gauss-Jordan on my own.
func (m Matrix4x4) Inverse() Matrix4x4 {
	indxc := []int{0, 0, 0, 0}
	indxr := []int{0, 0, 0, 0}
	ipiv := []int{0, 0, 0, 0}
	minv := make([][]float64, 4)
	for i, r := range m.M {
		minv[i] = make([]float64, 4)
		for j, v := range r {
			minv[i][j] = v
		}
	}

	for i := 0; i < 4; i++ {
		irow := -1
		icol := -1
		big := 0.
		// Choose pivot
		for j := 0; j < 4; j++ {
			if ipiv[j] != 1 {
				for k := 0; k < 4; k++ {
					if ipiv[k] == 0 {
						if math.Abs(minv[j][k]) >= big {
							big = (math.Abs(minv[j][k]))
							irow = j
							icol = k
						}
					} else if ipiv[k] > 1 {
						panic("Singular matrix in MatrixInvert")
					}
				}
			}
		}

		ipiv[icol]++
		// Swap rows _irow_ and _icol_ for pivot
		if irow != icol {
			for k := 0; k < 4; k++ {
				minv[irow][k], minv[icol][k] = minv[icol][k], minv[irow][k]
			}
		}
		indxr[i] = irow
		indxc[i] = icol
		if minv[icol][icol] == 0. {
			panic("Singular matrix in MatrixInvert")
		}

		// Set $m[icol][icol]$ to one by scaling row _icol_ appropriately
		pivinv := 1. / minv[icol][icol]
		minv[icol][icol] = 1.
		for j := 0; j < 4; j++ {
			minv[icol][j] *= pivinv
		}

		// Subtract this row from others to zero out their columns
		for j := 0; j < 4; j++ {
			if j != icol {
				save := minv[j][icol]
				minv[j][icol] = 0
				for k := 0; k < 4; k++ {
					minv[j][k] -= minv[icol][k] * save
				}
			}
		}
	}
	// Swap columns to reflect permutation
	for j := 3; j >= 0; j-- {
		if indxr[j] != indxc[j] {
			for k := 0; k < 4; k++ {
				minv[k][indxr[j]], minv[k][indxc[j]] = minv[k][indxc[j]], minv[k][indxr[j]]
			}
		}
	}
	return Matrix4x4{minv}
}
