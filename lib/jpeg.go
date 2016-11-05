package lib

import (
	_ "github.com/bronzdoc/slacky/lib/format"
	"image"
	"image/jpeg"
	"io"
	"os"
)

type JPEG struct{}

func (j *JPEG) Encode(w io.Writer, image image.Image) {
	jpeg.Encode(w, image, nil)
}

func (j *JPEG) Decode(file *os.File) (image.Image, error) {
	img, err := jpeg.Decode(file)
	return img, err
}
