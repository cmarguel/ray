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
			t.raytraceSample(samp)
		}
	}
}

func (t Task) raytraceSample(sample sampler.Sample) {
	ray := t.Camera.GenerateRay(sample)

	radiance := t.Renderer.Li(ray, t.World)

	t.Camera.Film().AddSample(sample, radiance)
}
