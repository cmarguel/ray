package accel

import (
	"container/list"
)

type Voxel struct {
	Primitives      *list.List
	AllCanIntersect bool
}

func NewVoxel(p Primitive) Voxel {
	primitives := list.New()
	primitives.PushBack(p)
	return Voxel{primitives, false}
}

func (v Voxel) AddPrimitive(p Primitive) {
	v.Primitives.PushBack(p)
}
