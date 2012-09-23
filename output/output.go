package output

import "image"

type Output interface {
	Output(image.Image)
}
