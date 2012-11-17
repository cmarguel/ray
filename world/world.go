package world

import (
	"ray/accel"
	"ray/geom"
	"ray/light"
	"ray/shape"
	"ray/visibility"
)

type World struct {
	Shapes    []shape.Shape
	Aggregate accel.Primitive
	Lights    []light.Light
}

func NewWorld() World {
	return World{make([]shape.Shape, 0), *new(accel.Primitive), make([]light.Light, 0)}
}

func (w *World) AddShape(shape shape.Shape) {
	w.Shapes = append(w.Shapes, shape)
}

func (w *World) SetPrimitive(aggregate accel.Primitive) {
	w.Aggregate = aggregate
}

func (w *World) AddLight(lite light.Light) {
	w.Lights = append(w.Lights, lite)
}

func (w *World) IntersectP(r geom.Ray) bool {
	return w.Aggregate.IntersectP(r)
}

func (w World) Unoccluded(t visibility.Tester) bool {
	return !w.IntersectP(t.R)
}
