package world

import (
	"ray/geom"
)

type World struct {
	Shapes []geom.Shape
}

func NewWorld() World {
	return World{make([]geom.Shape, 0)}
}

func (w *World) AddShape(shape geom.Shape) {
	w.Shapes = append(w.Shapes, shape)
}
