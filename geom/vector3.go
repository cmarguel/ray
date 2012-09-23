package geom

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (v1 Vector3) Dot(v2 Vector3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vector3) Plus(v2 Vector3) Vector3 {
	return Vector3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector3) Minus(v2 Vector3) Vector3 {
	return Vector3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v Vector3) Neg() Vector3 {
	return Vector3{-v.X, -v.Y, -v.Z}
}

func (v Vector3) MagnitudeSquared() float64 {
	return math.Pow(v.X, 2.) + math.Pow(v.Y, 2.) + math.Pow(v.Z, 2.)
}

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2.) + math.Pow(v.Y, 2.) + math.Pow(v.Z, 2.))
}

func (v Vector3) Normalized() Vector3 {
	return v.Times(1. / v.Magnitude())
}

func (v Vector3) IsZero() bool {
	return v.Magnitude() < 0.0000001
}

func (v Vector3) Times(a float64) Vector3 {
	return Vector3{a * v.X, a * v.Y, a * v.Z}
}
