package lib

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type Image struct {
	Width  uint
	Height uint
	Path   string
	image  image.Image
}

func NewImage(path string) *Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return &Image{
		Width:  0,
		Height: 0,
		Path:   path,
		image:  img,
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

	jpeg.Encode(out, resizedImage, nil)
}
