package render

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	//"math"

	"ray/camera"
	"ray/camera/film"
	"ray/integrator"
	"ray/output"
	"ray/render/sampler"
	"ray/world"
)

type Drawable interface {
	draw.Image
}

type Canvas struct {
	Width  int
	Height int

	image  Drawable
	output output.Output
	camera camera.Camera
	film   film.ImageFilm
}

func NewCanvasPNG(w, h int, filename string) Canvas {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	out := output.NewPNGOutput(filename)

	film := film.NewImageFilm(w, h)

	camera := camera.NewPinholeCamera(film)

	return Canvas{w, h, m, out, camera, film}
}

func (c Canvas) Set(x, y int, color color.Color) {
	c.Set(x, y, color)
}

func (c Canvas) Render(wor world.World) {
	c.render(wor)

	for i, pp := range c.film.Pixels {
		for j, p := range pp {
			r := uint8(p.Lxyz[0] * 255)
			g := uint8(p.Lxyz[1] * 255)
			b := uint8(p.Lxyz[2] * 255)
			c.image.Set(i, j, color.RGBA{r, g, b, 255})
		}
	}

	c.output.Output(c.image)
}

func (c Canvas) raytrace(x, y int, wor world.World, renderer sampler.Renderer) {
	sample := camera.NewCameraSample(x, y)
	ray := c.camera.GenerateRay(sample)

	radiance := renderer.Li(ray, wor)
	rf, gf, bf := radiance.ToRGB()

	c.film.Add(x, y, []float64{rf, gf, bf}, 1.)
}

func taskLogger(numTasks int, done <-chan int) {
	i := 0
	tenPercent := numTasks / 10
	for d := range done {
		i += d
		if i%tenPercent == 0 {
			fmt.Printf("%d%%\n", i*10/tenPercent)
		}
	}
}

func (c Canvas) render(wor world.World) {
	const numRoutines = 8
	runner := NewTaskRunner(numRoutines)
	runner.Start()

	whitted := integrator.NewWhitted()
	renderer := sampler.NewRenderer(whitted)

	totalRays := c.Width * c.Height
	numTasks := totalRays/(16*16) + 1
	go taskLogger(numTasks, runner.TasksDone)
	fmt.Printf("%d tasks to do with %d goroutines\n", numTasks, numRoutines)
	//onePercent := totalRays / 100
	for x := 0; x < c.Width+16; x += 16 {
		for y := 0; y < c.Height+16; y += 16 {
			task := NewTask(renderer, c, c.camera, wor, x, y)
			runner.AddTask(task)
		}
	}
	runner.Stop()
	runner.Wait()
}
