package geom

type Vector struct {
	p []float64
}

func NewVector2(x float64, y float64) Vector {
	return Vector{[]float64{x, y}}
}

func NewVector4(x, y, z, a float64) Vector {
	return Vector{[]float64{x, y, z, a}}
}

func (v Vector) Len() int {
	return len(v.p)
}

func (v1 Vector) Dot(v2 Vector) float64 {
	if v1.Len() != v2.Len() {
		panic("Tried to compute dot product of vectors with mismatching dimensions")
	}
	acc := 0.
	for i := range v1.p {
		acc += v1.p[i] * v2.p[i]
	}
	return acc
}
