package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"ray/output"
)

func NewCanvas(w int, h int) {
	image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
}

func main() {
	fmt.Println("Making basic image")

	m := NewCanvas(800, 600)
	m.Set(100, 100, color.RGBA{255, 0, 0, 255})

	out := output.NewPNGOutput("test.png")
	out.Output(m)

	fmt.Println("Done")

}
