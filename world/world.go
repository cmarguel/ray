package world

import (
	"ray/shape"
)

type World struct {
	Shapes []shape.Shape
}

func NewWorld() World {
	return World{make([]shape.Shape, 0)}
}

func (w *World) AddShape(shape shape.Shape) {
	w.Shapes = append(w.Shapes, shape)
}
