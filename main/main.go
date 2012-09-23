package main

import (
	"fmt"

	"ray/geom"
	"ray/render"
)

func makeRandomTriangle() geom.Triangle {
	white := geom.Color{255, 255, 255}
	p1 := geom.Vector3{200, 200, 200}
	p2 := geom.Vector3{200, 300, 200}
	p3 := geom.Vector3{300, 200, 200}

	v1 := geom.Vertex{p1, white}
	v2 := geom.Vertex{p2, white}
	v3 := geom.Vertex{p3, white}

	return geom.Triangle{v1, v2, v3}
}

func main() {
	fmt.Println("Making basic image")

	c := render.NewCanvasPNG(800, 600, "test.png")

	triangle := makeRandomTriangle()
	c.Render(triangle)

	fmt.Println("Done")

}
