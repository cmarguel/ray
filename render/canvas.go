package render

import (
	"image"
	"image/color"
	"image/draw"

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
}

func NewCanvasPNG(w, h int, filename string) Canvas {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	return Canvas{w, h, m, output.NewPNGOutput(filename)}
}

func (c Canvas) Set(x, y int, color color.Color) {
	c.Set(x, y, color)
}

func (c Canvas) Render(tri geom.Triangle) {
	Render(c.image, c.Width, c.Height, tri)

	c.output.Output(c.image)
}
