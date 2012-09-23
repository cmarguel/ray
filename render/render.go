package render

import (
	"image/color"
	"ray/geom"
)

func Render(img Drawable, w int, h int, triangle geom.Triangle) {
	eye := geom.Vector3{float64(w / 2), float64(h / 2), 0}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dir := geom.Vector3{float64(x), float64(y), 50.}
			ray := geom.Ray{eye, dir}

			i, status := ray.IntersectTriangle(triangle)
			if status != 1 {
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				c := uint8(i.Z)
				img.Set(x, y, color.RGBA{c, c, c, 255})
			}
		}
	}
}
