package geom

type Point Vector3
type Color Vector3

type Vertex struct {
	P Point
	C Color
}

type Triangle struct {
	V1 Vertex
	V2 Vertex
	V3 Vertex
}

func (t Triangle) Contains(p Point) bool {
	return false
}
