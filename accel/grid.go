package accel

import (
	"container/list"
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
					o := offset(x, y, z, g.nVoxels)
					if g.voxels[o] == nil {
						vox := NewVoxel(prim)
						g.voxels[o] = &vox
					} else {
						g.voxels[o].AddPrimitive(prim)
					}
				}
			}
		}

	}

}

func offset(x, y, z int, nVoxels []int) int {
	return z*nVoxels[0]*nVoxels[1] + y*nVoxels[0] + x
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
	if !intersected {
		return *new(Intersection), false
	}

	// Set up 3d dda for ray

	// walk ray through voxel grid

	return *new(Intersection), false
}

func (g Grid) setup3dDDA(ray *geom.Ray) {
	nextCrossingT := make([]float64, 3)
	deltaT := make([]float64, 3)
	rayD := ray.Direction.Minus(ray.Origin)

	step := make([]int, 3)
	out := make([]int, 3)
	pos := make([]int, 3)

	for axis := 0; axis < 3; axis++ {
		// compute current voxel for axis
		if rayD.Vals()[axis] >= 0 {
			// handle ray w/ pos direction 
		} else {
			// handle ray w/ neg direction
		}
	}
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
