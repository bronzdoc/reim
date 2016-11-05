package image

import (
	"github.com/bronzdoc/slacky/lib/format"
	"github.com/nfnt/resize"
	"image"
	"log"
	"os"
)

type Image struct {
	Width       uint
	Height      uint
	Path        string
	image       image.Image
	ImageFormat format.Formater
}

func NewImage(path string, imgf format.Formater) *Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := imgf.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return &Image{
		Width:       0,
		Height:      0,
		Path:        path,
		image:       img,
		ImageFormat: imgf,
	}
}

func (i *Image) Resize(width, height uint, name string) {
	i.Width = width
	i.Height = height

	resizedImage := resize.Resize(i.Width, i.Height, i.image, resize.Lanczos3)

	out, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	i.ImageFormat.Encode(out, resizedImage)
}
