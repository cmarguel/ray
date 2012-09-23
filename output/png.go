package output

import (
	"image"
	"image/png"
	"log"
	"os"
)

type PNGOutput struct {
	filename string
}

func (p PNGOutput) Output(img image.Image) {
	writeImageToFile(img, p.filename)
}

func NewPNGOutput(filename string) PNGOutput {
	return PNGOutput{filename}
}

func writeImageToFile(img image.Image, filename string) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
