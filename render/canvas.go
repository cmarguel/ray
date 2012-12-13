package render

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	//"math"

	"ray/camera"
	"ray/output"
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
}

func NewCanvasPNG(w, h int, filename string) Canvas {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	out := output.NewPNGOutput(filename)

	film := camera.Film{w, h}

	camera := camera.NewPinholeCamera(film)

	return Canvas{w, h, m, out, camera}
}

func (c Canvas) Set(x, y int, color color.Color) {
	c.Set(x, y, color)
}

func (c Canvas) Render(wor world.World) {
	c.render(wor)

	c.output.Output(c.image)
}

func (c Canvas) raytrace(x, y int, wor world.World) {
	sample := camera.NewCameraSample(x, y)
	ray := c.camera.GenerateRay(sample)

	// nearestPointDistance := math.Inf(1)
	intersect, found := wor.Aggregate.Intersect(&ray)
	radiance := evaluateRadiance(wor, &intersect.DiffGeom)
	rf, gf, bf := radiance.ToRGB()
	// fmt.Println(uint8(rf*255), uint8(gf*255), uint8(bf*255))
	col := color.RGBA{uint8(rf * 255), uint8(gf * 255), uint8(bf * 255), 255}
	c.image.Set(x, y, col)

	/*for _, shape := range wor.Shapes {
		dg, colFound, found := shape.Intersect(&ray)
		if found {
			distance := dg.P.DistanceSquared(c.camera.GetPos())
			if distance < nearestPointDistance {
				// fmt.Println("inside")
				//col := uint8(255. * (1 - distance/64.)) // uint8(255 - ((i.Z + 4) * 255 / 8))
				radiance := evaluateRadiance(wor, dg)
				rf, gf, bf := radiance.ToRGB()
				rf *= colFound.X // 255
				gf *= colFound.Y // 255
				bf *= colFound.Z // 255
				col := color.RGBA{uint8(rf), uint8(gf), uint8(bf), 255}
				// col := color.RGBA{uint8(colFound.X), uint8(colFound.Y), uint8(colFound.Z), 255}
				c.image.Set(x, y, col)
				nearestPointDistance = distance
			}
		}
	}*/
	// if math.IsInf(nearestPointDistance, 1) {
	if !found {
		c.image.Set(x, y, color.RGBA{0, 0, 0, 255})
	}
}

func (c Canvas) render(wor world.World) {
	runner := NewTaskRunner(4)
	runner.Start()

	totalRays := c.Width * c.Height
	onePercent := totalRays / 100
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			currentRays := x*c.Height + y
			if currentRays%(10*onePercent) == 0 {
				fmt.Printf("%d percent\n", currentRays/onePercent)
			}
			c.raytrace(x, y, wor)
		}
	}
}
