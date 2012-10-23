package accel

import (
	"container/list"
	"fmt"
	"math"
	"ray/geom"
	"ray/mmath"
)

type Grid struct {
	primitives *list.List
	bounds     geom.BBox
	nVoxels    []int
	voxels     []*Voxel
	invWidth   []float64
	width      []float64
}

func NewGrid(p *list.List, refineImmediately bool) Grid {
	grid := new(Grid)
	grid.Initialize(p, refineImmediately)

	return *grid
}

func (g *Grid) Initialize(p *list.List, refineImmediately bool) {
	g.primitives = nil
	g.bounds = geom.NewBBoxEmpty()
	g.nVoxels = make([]int, 3)
	g.width = make([]float64, 3)
	g.invWidth = make([]float64, 3)

	if refineImmediately {
		g.primitives = list.New()
		for e := p.Front(); e != nil; e = e.Next() {
			FullyRefine(e.Value.(Primitive), g.primitives)
		}
	} else {
		g.primitives = p
	}
	// compute bounds, choose grid res
	for e := p.Front(); e != nil; e = e.Next() {
		g.bounds = g.bounds.Union(e.Value.(Primitive).WorldBound())
	}
	fmt.Println("Max bounds: ", g.bounds.Max)
	fmt.Println("Min bounds: ", g.bounds.Min)
	delta := g.bounds.Max.Minus(g.bounds.Min)
	// find voxelsPerUnitDist
	maxAxis := g.bounds.MaximumExtent()
	invMaxWidth := 1. / delta.Vals()[maxAxis]
	cubeRoot := 3. * math.Pow(float64(g.primitives.Len()), 1./3.)
	voxelsPerUnitDist := cubeRoot * invMaxWidth

	for axis, del := range delta.Vals() {
		// This is a round function that will only work if the value is positive, which 
		// in this case should always be true
		g.nVoxels[axis] = int(math.Floor(del*voxelsPerUnitDist + 0.5))
		g.nVoxels[axis] = mmath.ClampInt(g.nVoxels[axis], 1, 64)
	}
	//g.nVoxels = []int{3, 3, 3}

	// compute voxel widths, allocate voxels
	for axis, del := range delta.Vals() {
		g.width[axis] = del / float64(g.nVoxels[axis])
		if g.width[axis] == 0. {
			g.invWidth[axis] = 0
		} else {
			g.invWidth[axis] = 1. / g.width[axis]
		}
	}
	nv := g.nVoxels[0] * g.nVoxels[1] * g.nVoxels[2]
	fmt.Println("voxel grid has dimensions: ", g.nVoxels[0], g.nVoxels[1], g.nVoxels[2])
	// In PBRT, this allocation was done via an aligned allocate, in order to be cache friendly.
	// I don't think we can do anything about that in go.
	g.voxels = make([]*Voxel, nv)
	for i := range g.voxels {
		g.voxels[i] = nil
	}

	// find voxel extent 
	for e := p.Front(); e != nil; e = e.Next() {
		prim := e.Value.(Primitive)
		pb := prim.WorldBound()
		vmin := make([]int, 3)
		vmax := make([]int, 3)

		for axis := 0; axis < 3; axis++ {
			vmin[axis] = g.posToVoxel(pb.Min, axis)
			vmax[axis] = g.posToVoxel(pb.Max, axis)
		}

		// add primitives
		for z := vmin[2]; z <= vmax[2]; z++ {
			for y := vmin[1]; y <= vmax[1]; y++ {
				for x := vmin[0]; x <= vmax[0]; x++ {
					o := g.offset(x, y, z)
					if g.voxels[o] == nil {
						vox := NewVoxel(prim)
						g.voxels[o] = &vox
						//fmt.Printf("(%d, %d, %d)\n", x, y, z)
					} else {
						g.voxels[o].AddPrimitive(prim)
						//fmt.Printf("(%d, %d, %d)\n", x, y, z)
					}
				}
			}
		}
	}
}

func (g Grid) offset(x, y, z int) int {
	return z*g.nVoxels[0]*g.nVoxels[1] + y*g.nVoxels[0] + x
}

func (g Grid) voxelToPos(p, axis int) float64 {
	return g.bounds.Min.Vals()[axis] + float64(p)*g.width[axis]
}

func (g Grid) posToVoxel(p geom.Vector3, axis int) int {
	v := int((p.Vals()[axis] - g.bounds.Min.Vals()[axis]) * g.invWidth[axis])
	return mmath.ClampInt(v, 0, g.nVoxels[axis]-1)
}

func (g Grid) WorldBound() geom.BBox {
	return *new(geom.BBox)
}

func (g Grid) CanIntersect() bool {
	return false
}

func (g Grid) Intersect(ray *geom.Ray) (Intersection, bool) {
	gridIntersect, rayT, intersected := g.checkRayAgainstGridBounds(ray)
	var intersect Intersection
	if !intersected {
		return intersect, false
	}

	// Set up 3d dda for ray
	nextCrossingT, deltaT, pos, step, out := g.setup3dDDA(ray, gridIntersect, rayT)

	// walk ray through voxel grid
	hitSomething := false
	for {
		// TODO mutex stuff here
		// check for intersection in current voxel, advance to next
		voxel := g.voxels[g.offset(pos[0], pos[1], pos[2])]
		if voxel != nil {
			hit := false
			intersect, hit = voxel.Intersect(ray)
			hitSomething = hitSomething || hit
		}

		stepAxis := g.computeStepAxis(nextCrossingT)

		if *ray.MaxT < nextCrossingT[stepAxis] {
			break
		}
		pos[stepAxis] += step[stepAxis]
		if pos[stepAxis] == out[stepAxis] {
			break
		}
		nextCrossingT[stepAxis] += deltaT[stepAxis]
	}

	return intersect, hitSomething
}

func (g Grid) computeStepAxis(next []float64) int {
	/*if next[0] >= next[1] && next[0] >= next[2] {
		return 0
	} else if next[1] >= next[2] && next[1] >= next[0] {
		return 1
	}
	return 2*/
	a := 0
	b := 0
	c := 0
	if next[0] < next[1] {
		a = 4
	}
	if next[0] < next[2] {
		b = 2
	}
	if next[1] < next[2] {
		c = 1
	}
	cmp := []int{2, 1, 2, 1, 2, 2, 0, 0}
	return cmp[a+b+c]
}

func (g Grid) setup3dDDA(ray *geom.Ray, gridIntersect geom.Vector3, rayT float64) ([]float64, []float64, []int, []int, []int) {
	nextCrossingT := make([]float64, 3)
	deltaT := make([]float64, 3)
	rayD := ray.Direction

	step := make([]int, 3)
	out := make([]int, 3)
	pos := make([]int, 3)

	for axis := 0; axis < 3; axis++ {
		// compute current voxel for axis
		pos[axis] = g.posToVoxel(gridIntersect, axis)

		if rayD.Vals()[axis] >= 0 {
			// handle ray w/ pos direction 
			nextCrossingT[axis] = rayT + (g.voxelToPos(pos[axis]+1, axis)-gridIntersect.Vals()[axis])/rayD.Vals()[axis]
			deltaT[axis] = g.width[axis] / rayD.Vals()[axis]
			step[axis] = 1
			out[axis] = g.nVoxels[axis]
		} else {
			// handle ray w/ neg direction
			nextCrossingT[axis] = rayT + (g.voxelToPos(pos[axis], axis)-gridIntersect.Vals()[axis])/rayD.Vals()[axis]
			deltaT[axis] = -g.width[axis] / rayD.Vals()[axis]
			step[axis] = -1
			out[axis] = -1
		}
	}

	return nextCrossingT, deltaT, pos, step, out
}

func (g Grid) checkRayAgainstGridBounds(ray *geom.Ray) (geom.Vector3, float64, bool) {
	rayT := 0.
	if g.bounds.Inside(ray.At(*ray.MinT)) {
		rayT = *ray.MinT
	} else {
		hit0, _, found := g.bounds.IntersectP(*ray)
		if !found {
			return *new(geom.Vector3), 0, false
		} else {
			rayT = hit0
		}
	}

	gridIntersect := ray.At(rayT)
	return gridIntersect, rayT, true
}

func (g Grid) Refine(*list.List) {
}
