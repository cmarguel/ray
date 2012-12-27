package shape

import (
	"container/list"
	"ray/geom"
)

type Mesh struct {
	Ind []int
	P   []geom.Vector3
}

func NewMesh(ind []int, p []geom.Vector3) Mesh {
	return Mesh{ind, p}
}

func (m Mesh) CanIntersect() bool {
	return false
}

func (m Mesh) Intersect(ray *geom.Ray) (*DifferentialGeometry, float64, float64, bool) {
	panic("Mesh's intersect method was called. The mesh should never be intersected; the refine method should have broken it down.")
	return nil, 0, 0, false
}

func (m Mesh) IntersectP(ray geom.Ray) bool {
	panic("Mesh's intersect method was called. The mesh should never be intersected; the refine method should have broken it down.")
	return false
}

func (m Mesh) WorldBound() geom.BBox {
	box := geom.NewBBoxEmpty()
	for _, v := range m.P {
		box = box.AddPoint(v)
	}
	return box
}

func (m Mesh) Refine(list *list.List) {
	for i := 0; i < len(m.Ind)-3; i++ {
		p1 := m.P[m.Ind[i]]
		p2 := m.P[m.Ind[i+1]]
		p3 := m.P[m.Ind[i+2]]
		list.PushBack(NewTriangle(p1.X, p1.Y, p1.Z, p2.X, p2.Y, p2.Z, p3.X, p3.Y, p3.Z))
	}
}
