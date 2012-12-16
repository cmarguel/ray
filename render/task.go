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

	sampler sampler.Sampler
}

func NewTask(renderer sampler.Renderer, canvas Canvas, cam camera.Camera,
	wor world.World, samp sampler.Sampler) Task {
	return Task{renderer, canvas, cam, wor, samp}
}

func (t Task) Run() {
	for samples, hasMore := t.sampler.GetMoreSamples(); hasMore; samples, hasMore = t.sampler.GetMoreSamples() {
		for _, samp := range samples {
			t.Canvas.raytrace(int(samp.ImageX), int(samp.ImageY), t.World, t.Renderer)
		}
	}
}
