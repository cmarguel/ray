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
}

func NewGrid(p *list.List, refineImmediately bool) Grid {
	// initialize primitives
	var primitives *list.List = nil
	bounds := geom.NewBBoxEmpty()
	nVoxels := make([]int, 3)
	width := make([]float64, 3)
	invWidth := make([]float64, 3)

	if refineImmediately {
		primitives = list.New()
		for e := p.Front(); e != nil; e = e.Next() {
			FullyRefine(e.Value.(Primitive), primitives)
		}
	} else {
		primitives = p
	}
	// compute bounds, choose grid res
	for e := p.Front(); e != nil; e = e.Next() {
		bounds = bounds.Union(e.Value.(Primitive).WorldBound())
	}
	delta := bounds.Max.Minus(bounds.Min)
	// find voxelsPerUnitDist
	maxAxis := bounds.MaximumExtent()
	invMaxWidth := 1. / delta.Vals()[maxAxis]
	cubeRoot := 3. * math.Pow(float64(primitives.Len()), 1./3.)
	voxelsPerUnitDist := cubeRoot * invMaxWidth

	for axis, del := range delta.Vals() {
		// This is a round function that will only work if the value is positive, which 
		// in this case should always be true
		nVoxels[axis] = int(math.Floor(del*voxelsPerUnitDist + 0.5))
		nVoxels[axis] = mmath.ClampInt(nVoxels[axis], 1, 64)
	}

	// compute voxel widths, allocate voxels
	for axis, del := range delta.Vals() {
		width[axis] = del / float64(nVoxels[axis])
		if width[axis] == 0. {
			invWidth[axis] = 0
		} else {
			invWidth[axis] = 1. / width[axis]
		}
	}
	nv := nVoxels[0] * nVoxels[1] * nVoxels[2]
	// In PBRT, this allocation was done via an aligned allocate, in order to be cache friendly.
	// I don't think we can do anything about that in go.
	voxels := make([]*Voxel, nv)
	for i := range voxels {
		voxels[i] = nil
	}

	// find voxel extent 
	for e := p.Front(); e != nil; e = e.Next() {
		prim := e.Value.(Primitive)
		pb := prim.WorldBound()
		vmin := make([]int, 3)
		vmax := make([]int, 3)

		for axis := 0; axis < 3; axis++ {
			vmin[axis] = posToVoxel(pb.Min, axis, bounds, invWidth, nVoxels)
			vmax[axis] = posToVoxel(pb.Max, axis, bounds, invWidth, nVoxels)
		}

		// add primitives
		for z := vmin[2]; z <= vmax[2]; z++ {
			for y := vmin[1]; y <= vmax[1]; y++ {
				for x := vmin[0]; x <= vmax[0]; x++ {
					o := offset(x, y, z, nVoxels)
					if voxels[o] == nil {
						vox := NewVoxel(prim)
						voxels[o] = &vox
					} else {
						voxels[o].AddPrimitive(prim)
					}
				}
			}
		}

	}

	// create reader-writer mutex

	return Grid{primitives, bounds, nVoxels, voxels}
}

func offset(x, y, z int, nVoxels []int) int {
	return z*nVoxels[0]*nVoxels[1] + y*nVoxels[0] + x
}

func posToVoxel(p geom.Vector3, axis int, bounds geom.BBox, invWidth []float64, nVoxels []int) int {
	v := int((p.Vals()[axis] - bounds.Min.Vals()[axis]) * invWidth[axis])
	return mmath.ClampInt(v, 0, nVoxels[axis]-1)
}

func (g Grid) WorldBound() geom.BBox {
	return *new(geom.BBox)
}

func (g Grid) CanIntersect() bool {
	return false
}

func (g Grid) Intersect(*geom.Ray) (Intersection, bool) {
	return *new(Intersection), false
}

func (g Grid) Refine(*list.List) {
}
