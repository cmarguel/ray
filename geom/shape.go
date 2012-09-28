package geom

type Shape interface {
	Intersect(ray Ray) (Vector3, bool)
}
