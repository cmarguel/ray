package shape

import (
	"math"
	"ray/geom"
	"ray/mmath"
)

type Cube struct {
	triangles []Triangle
}

func NewCube() Cube {
	side11 := NewTriangle(
		-1, 1, -1,
		-1, -1, -1,
		1, 1, -1)
	side12 := NewTriangle(
		-1, -1, -1,
		1, -1, -1,
		1, 1, -1)

	side21 := NewTriangle(
		1, 1, -1,
		1, -1, -1,
		1, 1, 1)
	side22 := NewTriangle(
		1, -1, -1,
		1, -1, 1,
		1, 1, 1)

	side31 := NewTriangle(
		1, 1, 1,
		1, -1, 1,
		-1, -1, 1)
	side32 := NewTriangle(
		1, 1, 1,
		-1, -1, 1,
		-1, 1, 1)

	side41 := NewTriangle(
		-1, 1, 1,
		-1, -1, 1,
		-1, -1, -1)
	side42 := NewTriangle(
		-1, 1, 1,
		-1, -1, -1,
		-1, 1, -1)

	side51 := NewTriangle(
		1, 1, 1,
		-1, 1, -1,
		-1, 1, 1)
	side52 := NewTriangle(
		1, 1, 1,
		1, 1, -1,
		-1, 1, -1)

	side61 := NewTriangle(
		-1, -1, 1,
		-1, -1, -1,
		1, -1, 1)
	side62 := NewTriangle(
		-1, -1, -1,
		1, -1, -1,
		1, -1, 1)

	triangles := []Triangle{side11, side12, side21, side22, side31, side32, side41, side42,
		side51, side52, side61, side62}
	return Cube{triangles}
}

func (c Cube) Intersect(ray geom.Ray) (geom.Vector3, bool) {
	nearest := math.Inf(1)
	intersection := geom.NewVector3(0, 0, 0)
	for _, t := range c.triangles {
		i, found := t.Intersect(ray)
		if found {
			distance := i.DistanceSquared(ray.Origin)
			if distance < nearest {
				nearest = distance
				intersection = i
			}
		}
	}
	return intersection, !math.IsInf(nearest, 1)
}

func (c Cube) Transform(transform mmath.Transform) Cube {
	newTriangles := make([]Triangle, 0, 12)
	for _, t := range c.triangles {
		tr := t.Transform(transform)
		newTriangles = append(newTriangles, tr)
	}
	c.triangles = newTriangles
	return c
}
