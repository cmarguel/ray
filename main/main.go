package main

import (
	"fmt"
	"math"
	"math/rand"

	"ray/geom"
	"ray/mmath"
	"ray/render"
)

func randomlyOrientedTriangle() geom.Triangle {
	scale := rand.Float64() + 0.5
	rx := rand.Float64() * 2. * math.Pi
	ry := rand.Float64() * 2. * math.Pi
	rz := rand.Float64() * 2. * math.Pi
	dx := rand.Float64()*5. + 2.
	dy := rand.Float64()*5. + 2.
	dz := rand.Float64()*5. + 2.

	transform := mmath.NewTransform().
		Scale(scale, scale, scale).
		RotateX(rx).
		RotateY(ry).
		RotateZ(rz).
		Translate(dx, dy, dz)

	triangle := makeStaticTriangle()
	triangle.V1.P = transform.Apply(triangle.V1.P)
	triangle.V2.P = transform.Apply(triangle.V2.P)
	triangle.V3.P = transform.Apply(triangle.V3.P)

	return triangle
}

func makeStaticTriangle() geom.Triangle {
	white := geom.Color{255, 255, 255}
	p1 := geom.Vector3{0., 0., 0.}
	p2 := geom.Vector3{0., 1., 0.}
	p3 := geom.Vector3{1., 0., 0.}

	v1 := geom.Vertex{p1, white}
	v2 := geom.Vertex{p2, white}
	v3 := geom.Vertex{p3, white}

	return geom.Triangle{v1, v2, v3}
}

func main() {
	fmt.Println("Making basic image")

	c := render.NewCanvasPNG(800, 600, "test.png")

	const numTriangles = 10

	triangles := make([]geom.Triangle, 0, numTriangles)
	for i := 0; i < numTriangles; i++ {
		t := randomlyOrientedTriangle()
		triangles = append(triangles, t)
	}

	c.Render(triangles)

	fmt.Println("Done")

}
