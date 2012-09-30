package accel

import (
	"container/list"
)

type Grid struct {
	primitives *list.List
}

func NewGrid(p *list.List, refineImmediately bool) Grid {
	// initialize primitives
	var primitives *list.List = nil

	if refineImmediately {
		primitives = list.New()
		for e := p.Front(); e != nil; e = e.Next() {
			FullyRefine(e.Value.(Primitive), primitives)
		}
	} else {
		primitives = p
	}
	// compute bounds, choose grid res
	// compute voxel widths, allocate voxels
	// add primitives
	// create reader-writer mutex

	return *new(Grid)
}
