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
	"ray/mmath"
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

func NewCanvasPNG(w, h int, cam camera.Camera, filename string) Canvas {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	out := output.NewPNGOutput(filename)

	return Canvas{w, h, m, out, cam, cam.Film().(film.ImageFilm)}
}

func (c Canvas) Set(x, y int, color color.Color) {
	c.Set(x, y, color)
}

func (c Canvas) Render(wor world.World) {
	c.render(wor)

	for i, pp := range c.film.Pixels {
		for j, p := range pp {
			sum := p.WeightSum
			r := uint8((p.Lxyz[0] / sum) * 255)
			g := uint8((p.Lxyz[1] / sum) * 255)
			b := uint8((p.Lxyz[2] / sum) * 255)
			c.image.Set(i, j, color.RGBA{r, g, b, 255})
		}
	}

	c.output.Output(c.image)
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
	const numRoutines = 4
	runner := NewTaskRunner(numRoutines)
	runner.Start()

	whitted := integrator.NewWhitted()
	renderer := sampler.NewRenderer(whitted)

	numPixels := c.Width * c.Height
	numTasks := mmath.RoundUpPow2(numPixels / (16 * 16))
	go taskLogger(numTasks, runner.TasksDone)
	fmt.Printf("%d tasks to do with %d goroutines\n", numTasks, numRoutines)

	//mainSampler := sampler.NewStratified(0, c.Width, 0, c.Height, 4, 4, true, 0, 0)
	mainSampler := sampler.NewUniformSampler(0, c.Width, 0, c.Height, 1, 0, 0)

	for i := 0; i < numTasks; i++ {
		subsampler := mainSampler.GetSubSampler(numTasks-1-i, numTasks)
		task := NewTask(renderer, c, c.camera, wor, subsampler)
		runner.AddTask(task)
	}
	runner.Stop()
	runner.Wait()
}
