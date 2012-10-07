package world

import (
	"ray/accel"
	"ray/light"
	"ray/shape"
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
