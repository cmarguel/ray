package render

import (
	//"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"ray/camera"
	"ray/geom"
	"ray/output"
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

func (c Canvas) Render(tri []geom.Triangle) {
	c.render(tri)

	c.output.Output(c.image)
}

func (c Canvas) raytrace(x, y int, triangles []geom.Triangle) {
	sample := camera.NewCameraSample(x, y)
	ray := c.camera.GenerateRay(sample)

	nearestPointDistance := math.Inf(1)
	for _, tri := range triangles {
		i, status := ray.IntersectTriangle(tri)
		if status == 1 {
			distance := i.DistanceSquared(c.camera.GetPos())
			if distance < nearestPointDistance {
				// fmt.Println("inside")
				col := uint8(255. * (1 - distance/64.)) // uint8(255 - ((i.Z + 4) * 255 / 8))
				c.image.Set(x, y, color.RGBA{col, col, col, 255})
				nearestPointDistance = distance
			}
		}
	}
	if math.IsInf(nearestPointDistance, 1) {
		c.image.Set(x, y, color.RGBA{0, 0, 0, 255})
	}
}

func (c Canvas) render(triangles []geom.Triangle) {
	//totalRays := c.Width*c.Height
	//onePercent := totalRays / 100
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			//currentRays := x*c.Height + y
			//if currentRays % (5*onePercent) == 0 {
			//	fmt.Printf("%d percent\n", currentRays/onePercent)
			//} 
			c.raytrace(x, y, triangles)
		}
	}
}
