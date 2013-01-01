package main

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

	"ray/accel"
	"ray/camera"
	"ray/camera/film"
	"ray/geom"
	"ray/light"
	"ray/mmath"
	"ray/parser"
	"ray/render"
	"ray/render/sampler"
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
				RotateY(4*math.Pi/36).
				Translate(x-1, height/2.-8.5, z+3)
			cube := shape.NewCube()
			cubes = append(cubes, cube.Transform(baseTransform))
		}
	}
	return cubes
}

func main() {
	fmt.Println("Making basic image")

	conf := parser.LoadConfig("scenes/cornell-mlt.pbrt")
	//conf := parser.LoadConfig("scenes/microcity.pbrt")

	cam := setupCamera(800, 600)
	c := render.NewCanvasPNG(800, 600, cam, "test.png")
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

	cubes := makeGrid(numCubes, numCubes*3)
	shapeList := list.New()
	for _, c := range cubes {
		_ = c
		//wor.AddShape(c)
		//shapeList.PushBack(accel.NewGeometricPrimitive(c))
	}

	for _, att := range conf.Attributes {
		for _, sh := range att.Shapes {
			_ = sh
			sh = sh.(shape.Mesh).Transform(mmath.NewTransform().
				Scale(1./60., 1./60., 1./60.).
				RotateY(-5*math.Pi/36.).
				Translate(-6, -5.5, 12))
			shapeList.PushBack(accel.NewGeometricPrimitive(sh))
		}
	}

	fmt.Println("Building grid...")
	grid := accel.NewGrid(shapeList, false)
	wor.SetPrimitive(grid)
	fmt.Println("Done building grid!")
	fmt.Println("Bounding box for grid: ", grid.WorldBound())

	//cube3 := shape.NewCube()
	//tr := mmath.NewTransform().
	//	Scale(0.15, 0.15, 0.15).
	//	Translate(0, 0.5, 7.)

	//cube3 = cube3.Transform(tr)
	//wor.AddShape(cube3)

	//wor.AddLight(light.NewPointLight(1., 6, 13., 20., 20., 20.))
	wor.AddLight(light.NewPointLight(5., 0, 2., 120., 120., 120.))
	//wor.AddLight(light.NewPointLight(0., 0, 0., 120., 120., 120.))
	//wor.AddLight(light.NewPointLight(278., 335, 279.5, 100., 100., 100.))

	c.Render(wor)

	fmt.Println("Done")

}

func setupCamera(w, h int) camera.Camera {
	filter := sampler.NewGaussianFilter(0.5, 0.5, 0.75)
	film := film.NewImageFilm(w, h, filter)

	//camera := camera.NewPinholeCamera(film)
	frame := float64(film.ResolutionX) / float64(film.ResolutionY)
	screen := []float64{0, 0, 0, 0}
	if frame > 1. {
		screen = []float64{-frame, frame, -1., 1.}
	} else {
		screen = []float64{-1., 1., -1. / frame, 1. / frame}
	}

	c2w := mmath.LookAt(
		geom.NewVector3(0, 0, 0),
		geom.NewVector3(0.25, 0, 1),
		geom.NewVector3(0, 1, 0))
	//c2w = mmath.LookAt(
	//	geom.NewVector3(-1.42702, -3.30238, 1.79759),
	//	geom.NewVector3(0.023598, 9.69691, -4.68208),
	//	geom.NewVector3(0.00016145, 0.388419, 0.921483))

	_ = c2w
	_ = screen
	cam := camera.NewPerspective(c2w, screen, 55., film)
	//cam := camera.NewPinholeCamera(film)
	return cam
}
