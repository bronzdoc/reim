package format

import (
	"image"
	_png "image/png"
	"io"
	"os"
)

type png struct {
	magicNumbers []string
	extension    string
}

func NewPng() *png {
	return &png{
		magicNumbers: []string{"\x89PNG\r\n\x1a\n"},
		extension:    ".png",
	}
}

func (p *png) Encode(w io.Writer, image image.Image) {
	_png.Encode(w, image)
}

func (p *png) Decode(file *os.File) (image.Image, error) {
	img, err := _png.Decode(file)
	return img, err
}

func (p *png) MagicNumbers() []string {
	return p.magicNumbers
}

func (p *png) Extension() string {
	return p.extension
}
