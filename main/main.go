package main

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

	"ray/accel"
	"ray/geom"
	"ray/light"
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
	// white := geom.Color{255, 255, 255}
	p1 := geom.Vector3{0., 0., 0.}
	p2 := geom.Vector3{0., 1., 0.}
	p3 := geom.Vector3{1., 0., 0.}

	//v1 := geom.Vertex{p1, white}
	//v2 := geom.Vertex{p2, white}
	//v3 := geom.Vertex{p3, white}

	return shape.Triangle{p1, p2, p3}
}

func makeGrid(r, c int) []shape.Cube {
	const highestX = 5.
	const highestZ = 5.

	cubes := make([]shape.Cube, 0, r*c)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			//z := float64(j)*0.5 + 0.2
			//x := float64(i)*0.5 - float64(r/2)
			x := float64(i)*0.75 - float64(r)/2.
			z := float64(j)*0.75 + 0.75

			height := math.Abs(z-highestZ) + math.Abs(x-highestX)

			baseTransform := mmath.NewTransform().
				Scale(0.25, height*0.25, 0.25).
				Translate(x, 8.5-height/2., z)
			cube := shape.NewCube()
			cubes = append(cubes, cube.Transform(baseTransform))
		}
	}
	return cubes
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
		// t := randomlyOrientedTriangle()
		// wor.AddShape(t)
	}

	numCubes := 10
	if len(os.Args) > 2 {
		numCubes, _ = strconv.Atoi(os.Args[2])
	}

	/* for i := 0; i < numCubes; i++ {
		c := shape.NewCube().Transform(randomTransform())

		wor.AddShape(c)
	}*/

	cubes := makeGrid(numCubes, numCubes)
	cubeList := list.New()
	for _, c := range cubes {
		// wor.AddShape(c)
		cubeList.PushBack(accel.NewGeometricPrimitive(c))
	}
	fmt.Println("Building grid...")
	grid := accel.NewGrid(cubeList, false)
	wor.SetPrimitive(grid)
	fmt.Println("Done building grid!")

	//cube3 := shape.NewCube()
	//tr := mmath.NewTransform().
	//	Scale(0.15, 0.15, 0.15).
	//	Translate(0, 0.5, 7.)

	//cube3 = cube3.Transform(tr)
	//wor.AddShape(cube3)

	wor.AddLight(light.NewPointLight(1., -2.5, 4., 10., 10., 10.))

	c.Render(wor)

	fmt.Println("Done")

}
