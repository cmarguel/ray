package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"ray/geom"
	"ray/output"
)

type Drawable interface {
	draw.Image
}

func NewCanvas(w int, h int) Drawable {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	return m
}

func Render(img Drawable, w int, h int, triangle geom.Triangle) {
	eye := geom.Vector3{float64(w / 2), float64(h / 2), 0}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dir := geom.Vector3{float64(x), float64(y), 50.}
			ray := geom.Ray{eye, dir}

			_, status := ray.IntersectTriangle(triangle)
			if status != 1 {
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				c := uint8(triangle.V1.P.Z)
				img.Set(x, y, color.RGBA{c, c, c, 255})
			}
		}
	}
}

func main() {
	fmt.Println("Making basic image")

	m := NewCanvas(800, 600)
	m.Set(100, 100, color.RGBA{255, 0, 0, 255})

	white := geom.Color{255, 255, 255}
	p1 := geom.Vector3{200, 200, 200}
	p2 := geom.Vector3{200, 300, 200}
	p3 := geom.Vector3{300, 200, 200}

	v1 := geom.Vertex{p1, white}
	v2 := geom.Vertex{p2, white}
	v3 := geom.Vertex{p3, white}

	triangle := geom.Triangle{v1, v2, v3}

	Render(m, 800, 600, triangle)

	out := output.NewPNGOutput("test.png")
	out.Output(m)

	fmt.Println("Done")

}
