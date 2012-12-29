package mmath

import (
	"fmt"
	"math"
	"testing"
)

func Test_inverse_from_transform(t *testing.T) {
	m := NewTransform().Scale(1, 2, 3).RotateX(3.14159).Translate(4, 5, 6).m
	inv := m.Inverse()
	ident := NewMatrix4x4(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)

	prod := m.Times(inv)
	if !matEq(prod, ident) {
		fmt.Printf("Matrix was not inverted properly. Was: %s\n", prod)
	}
}

func Test_inverse_from_perspective(t *testing.T) {
	m := NewTransform().Perspective(55., 1e-3, 1000.).m
	inv := m.Inverse()
	ident := NewMatrix4x4(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)

	prod := m.Times(inv)
	if !matEq(prod, ident) {
		fmt.Printf("Matrix was not inverted properly. Was: %s\n", prod)
	}
}

func Test_inverse_from_random(t *testing.T) {
	m := NewMatrix4x4(1, 2, 3, 3, 3, 2, 1, 2, 3, 1, 2, 1, 1, 1, 2, 2)
	inv := m.Inverse()
	ident := NewMatrix4x4(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)

	prod := m.Times(inv)
	if !matEq(prod, ident) {
		fmt.Printf("Matrix was not inverted properly. Product was: %s\nInverse was: %s\n", prod, inv)
	}
}

func matEq(m1, m2 Matrix4x4) bool {
	a := m1.M
	b := m2.M

	for i, r := range a {
		for j, v := range r {
			if math.Abs(v-b[i][j]) > 1e-5 {
				return false
			}
		}
	}
	return true
}
