package math

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

func (mat Matrix4x4) Row(i int) []float64 {
	m := mat.M
	return []float64{m[i][0], m[i][1], m[i][2], m[i][3]}
}

func (mat Matrix4x4) Col(i int) []float64 {
	m := mat.M
	return []float64{m[0][i], m[1][i], m[2][i], m[3][i]}
}

func dot(v1, v2 Matrix4x4, col, row int) float64 {
	a := v1.Col(col)
	b := v2.Row(row)
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

func (m1 Matrix4x4) Times(m2 Matrix4x4) Matrix4x4 {
	t00 := dot(m1, m2, 0, 0)
	t01 := dot(m1, m2, 0, 1)
	t02 := dot(m1, m2, 0, 2)
	t03 := dot(m1, m2, 0, 3)

	t10 := dot(m1, m2, 1, 0)
	t11 := dot(m1, m2, 1, 1)
	t12 := dot(m1, m2, 1, 2)
	t13 := dot(m1, m2, 1, 3)

	t20 := dot(m1, m2, 2, 0)
	t21 := dot(m1, m2, 2, 1)
	t22 := dot(m1, m2, 2, 2)
	t23 := dot(m1, m2, 2, 3)

	t30 := dot(m1, m2, 3, 0)
	t31 := dot(m1, m2, 3, 1)
	t32 := dot(m1, m2, 3, 2)
	t33 := dot(m1, m2, 3, 3)

	return Matrix4x4{
		[][]float64{
			[]float64{t00, t01, t02, t03},
			[]float64{t10, t11, t12, t13},
			[]float64{t20, t21, t22, t23},
			[]float64{t30, t31, t32, t33},
		},
	}
}
