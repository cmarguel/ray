package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

	"ray/geom"
	"ray/mmath"
	"ray/render"
	"ray/shape"
	"ray/world"
)

func randomTransform() mmath.Transform {
	scale := rand.Float64() + 0.5
	rx := rand.Float64() * 2. * math.Pi
	ry := rand.Float64() * 2. * math.Pi
	rz := rand.Float64() * 2. * math.Pi
	dx := rand.Float64()*8. - 4.
	dy := rand.Float64()*8. - 4.
	dz := rand.Float64()*4 + 4.

	return mmath.NewTransform().
		Scale(scale, scale, scale).
		RotateX(rx).
		RotateY(ry).
		RotateZ(rz).
		Translate(dx, dy, dz)

}

func randomlyOrientedTriangle() shape.Triangle {
	transform := randomTransform()

	triangle := makeStaticTriangle()
	tr := triangle.Transform(transform)
	return tr
}

func makeStaticTriangle() shape.Triangle {
	white := geom.Color{255, 255, 255}
	p1 := geom.Vector3{0., 0., 0.}
	p2 := geom.Vector3{0., 1., 0.}
	p3 := geom.Vector3{1., 0., 0.}

	v1 := geom.Vertex{p1, white}
	v2 := geom.Vertex{p2, white}
	v3 := geom.Vertex{p3, white}

	return shape.Triangle{v1, v2, v3, geom.Color{255, 255, 255}}
}

func main() {
	fmt.Println("Making basic image")

	c := render.NewCanvasPNG(800, 600, "test.png")
	wor := world.NewWorld()

	numTriangles := 10
	if len(os.Args) > 1 {
		numTriangles, _ = strconv.Atoi(os.Args[1])
	}

	for i := 0; i < numTriangles; i++ {
		t := randomlyOrientedTriangle()
		wor.AddShape(t)
	}

	numCubes := 10
	if len(os.Args) > 2 {
		numTriangles, _ = strconv.Atoi(os.Args[2])
	}

	for i := 0; i < numCubes; i++ {
		c := shape.NewCube().Transform(randomTransform())
		wor.AddShape(c)
	}

	//cube3 := shape.NewCube()
	//tr := mmath.NewTransform().
	//	Scale(0.15, 0.15, 0.15).
	//	Translate(0, 0.5, 7.)

	//cube3 = cube3.Transform(tr)
	//wor.AddShape(cube3)

	c.Render(wor)

	fmt.Println("Done")

}
