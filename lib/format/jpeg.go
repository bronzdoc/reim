package format

import (
	"image"
	_jpeg "image/jpeg"
	"io"
	"os"
)

type jpeg struct {
	magicNumber string
}

func NewJpeg() *jpeg {
	return &jpeg{magicNumber: "\xff\xd8\xff"}
}

func (j *jpeg) Encode(w io.Writer, image image.Image) {
	_jpeg.Encode(w, image, nil)
}

func (j *jpeg) Decode(file *os.File) (image.Image, error) {
	img, err := _jpeg.Decode(file)
	return img, err
}

func (j *jpeg) MagicNumber() string {
	return j.magicNumber
}
