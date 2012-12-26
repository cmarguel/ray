package shape

import (
	"container/list"
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

	//green := geom.Color{0, 255, 0}
	//side11.Color = green
	//side12.Color = green

	side21 := NewTriangle(
		1, 1, -1,
		1, -1, -1,
		1, 1, 1)
	side22 := NewTriangle(
		1, -1, -1,
		1, -1, 1,
		1, 1, 1)

	//red := geom.Color{255, 0, 0}
	//side21.Color = red
	//side22.Color = red

	side31 := NewTriangle(
		1, 1, 1,
		1, -1, 1,
		-1, -1, 1)
	side32 := NewTriangle(
		1, 1, 1,
		-1, -1, 1,
		-1, 1, 1)

	//blue := geom.Color{0, 0, 255}
	//side31.Color = blue
	//side32.Color = blue

	side41 := NewTriangle(
		-1, 1, 1,
		-1, -1, 1,
		-1, -1, -1)
	side42 := NewTriangle(
		-1, 1, 1,
		-1, -1, -1,
		-1, 1, -1)

	//orange := geom.Color{255, 128, 0}
	//side41.Color = orange
	//side42.Color = orange

	side51 := NewTriangle(
		1, 1, 1,
		-1, 1, -1,
		-1, 1, 1)
	side52 := NewTriangle(
		1, 1, 1,
		1, 1, -1,
		-1, 1, -1)

	//white := geom.Color{255, 255, 255}
	//side51.Color = white
	//side52.Color = white

	side61 := NewTriangle(
		-1, -1, 1,
		-1, -1, -1,
		1, -1, 1)
	side62 := NewTriangle(
		-1, -1, -1,
		1, -1, -1,
		1, -1, 1)

	//yellow := geom.Color{255, 255, 0}
	//side61.Color = yellow
	//side62.Color = yellow

	triangles := []Triangle{side11, side12, side21, side22, side31, side32, side41, side42,
		side51, side52, side61, side62}
	return Cube{triangles}
}

func (c Cube) Intersect(ray *geom.Ray) (*DifferentialGeometry, float64, float64, bool) {
	nearest := math.Inf(1)
	var diffGeom *DifferentialGeometry = nil
	tHit := 0.
	for _, t := range c.triangles {
		dg, tH, _, found := t.Intersect(ray)
		if found {
			distance := tH
			if distance < nearest {
				nearest = distance
				diffGeom = dg
				tHit = tH
			}
		}
	}

	return diffGeom, tHit, 0.001, !math.IsInf(nearest, 1)
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

func (c Cube) WorldBound() geom.BBox {
	box := geom.NewBBoxEmpty()
	for _, t := range c.triangles {
		box = box.Union(t.WorldBound())
	}
	return box
}

func (c Cube) Refine(l *list.List) {
	for _, t := range c.triangles {
		t.Refine(l)
	}
}

func (c Cube) IntersectP(ray geom.Ray) bool {
	for _, t := range c.triangles {
		if t.IntersectP(ray) {
			return true
		}
	}
	return false
}

func (c Cube) CanIntersect() bool {
	return true
}
