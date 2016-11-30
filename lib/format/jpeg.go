package format

import (
	"image"
	_jpeg "image/jpeg"
	"io"
	"os"
)

type jpeg struct {
	magicNumbers []string
	extension    string
}

func NewJpeg() *jpeg {
	return &jpeg{
		magicNumbers: []string{"\xff\xd8\xff"},
		extension:    ".jpeg",
	}
}

func (j *jpeg) Encode(w io.Writer, image image.Image) {
	_jpeg.Encode(w, image, nil)
}

func (j *jpeg) Decode(file *os.File) (image.Image, error) {
	img, err := _jpeg.Decode(file)
	return img, err
}

func (j *jpeg) MagicNumbers() []string {
	return j.magicNumbers
}
func (j *jpeg) Extension() string {
	return j.extension
}
