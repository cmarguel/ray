package render

import (
	//"fmt"
	"ray/camera"
	"ray/render/sampler"
	"ray/world"
)

type Task struct {
	Renderer sampler.Renderer
	Canvas   Canvas
	Camera   camera.Camera
	World    world.World
	MinX     int
	MinY     int
}

func NewTask(renderer sampler.Renderer, canvas Canvas, cam camera.Camera,
	wor world.World, minX, minY int) Task {
	return Task{renderer, canvas, cam, wor, minX, minY}
}

func (t Task) Run() {
	for x := t.MinX; x < t.MinX+16; x++ {
		for y := t.MinY; y < t.MinY+16; y++ {
			t.Canvas.raytrace(x, y, t.World, t.Renderer)
		}
	}
}
