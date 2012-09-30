package accel

import (
	"container/list"
	"math"
	"ray/mmath"
)

type Grid struct {
	primitives *list.List
	bounds     BBox
	nVoxels    []int
}

func NewGrid(p *list.List, refineImmediately bool) Grid {
	// initialize primitives
	var primitives *list.List = nil
	bounds := NewBBoxEmpty()
	nVoxels := make([]int, 3)

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
	voxelsPerUnitDist := 0.

	for axis, del := range delta.Vals() {
		// This is a round function that will only work if the value is positive, which 
		// in this case should always be true
		nVoxels[axis] = int(math.Floor(del*voxelsPerUnitDist + 0.5))
		nVoxels[axis] = mmath.ClampInt(nVoxels[axis], 1, 64)
	}

	// compute voxel widths, allocate voxels
	// add primitives
	// create reader-writer mutex

	return *new(Grid)
}
