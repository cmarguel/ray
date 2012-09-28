package geom

type Point Vector3
type Color Vector3

type Vertex struct {
	P Vector3
	C Color
}

type Ray struct {
	Origin    Vector3
	Direction Vector3
}
