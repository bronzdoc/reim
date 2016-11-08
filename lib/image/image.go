package image

import (
	"fmt"
	"github.com/bronzdoc/reim/lib/format"
	"github.com/nfnt/resize"
	"image"
	"log"
	"os"
	"time"
)

type Image struct {
	Width       uint
	Height      uint
	Path        string
	image       image.Image
	ImageFormat format.Formater
	name        string
}

func NewImage(path string) *Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	imgf, err := format.NewFormatManager(path).Build()
	if err != nil {
		log.Fatal(err)
	}

	img, err := imgf.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return &Image{
		Path:        path,
		image:       img,
		ImageFormat: imgf,
		name: fmt.Sprint(
			fmt.Sprint(time.Now().Unix()),
			imgf.Extension(),
		),
	}
}

func (i *Image) Resize(width, height uint) {
	i.Width = width
	i.Height = height

	resizedImage := resize.Resize(i.Width, i.Height, i.image, resize.Lanczos3)

	out, err := os.Create(i.name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	i.ImageFormat.Encode(out, resizedImage)
}

func (i *Image) Rename(name string) {
	i.name = name
}

func (i *Image) Name() string {
	return i.name
}
