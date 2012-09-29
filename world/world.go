package world

import (
	"ray/light"
	"ray/shape"
)

type World struct {
	Shapes []shape.Shape
	Lights []light.Light
}

func NewWorld() World {
	return World{make([]shape.Shape, 0), make([]light.Light, 0)}
}

func (w *World) AddShape(shape shape.Shape) {
	w.Shapes = append(w.Shapes, shape)
}

func (w *World) AddLight(lite light.Light) {
	w.Lights = append(w.Lights, lite)
}
