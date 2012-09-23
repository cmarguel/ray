package main

import (
	"fmt"
	//"math"
	//"math/rand"

	"ray/geom"
	"ray/render"
)

func randomlyOrientedTriangle() {
	// rand.Float64() * 2. * math.Pi

	// triangle
}

func makeStaticTriangle() geom.Triangle {
	white := geom.Color{255, 255, 255}
	p1 := geom.Vector3{-1., 1., 3.}
	p2 := geom.Vector3{-1., 2., 3.}
	p3 := geom.Vector3{-2., 1., 3.}

	v1 := geom.Vertex{p1, white}
	v2 := geom.Vertex{p2, white}
	v3 := geom.Vertex{p3, white}

	return geom.Triangle{v1, v2, v3}
}

func main() {
	fmt.Println("Making basic image")

	c := render.NewCanvasPNG(800, 600, "test.png")

	triangles := []geom.Triangle{makeStaticTriangle()}
	c.Render(triangles)

	fmt.Println("Done")

}
