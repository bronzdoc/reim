package main

import (
	"flag"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type Options struct {
	Width     uint
	Height    uint
	Out       string
	ImageName string
}

type Image struct {
	Width  uint
	Height uint
	Path   string
	image  image.Image
}

func newImage(path string) *Image {
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

func (i *Image) resize(width, height uint, name string) {
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

type Cli struct {
	Options Options
}

func newCli(options Options) *Cli {
	return &Cli{Options: options}
}

func (c *Cli) run(image *Image) {
	width := c.Options.Width
	height := c.Options.Height
	image.resize(width, height, "slack-image.jpg")
}

var (
	cli *Cli
)

func init() {
	options := Options{}
	flag.StringVar(&options.ImageName, "image", "", "Name of the image to be resized")
	flag.UintVar(&options.Width, "width", 20, "New image width")
	flag.UintVar(&options.Height, "height", 20, "New image height")
	flag.StringVar(&options.Out, "out", "slack-image.jpg", "new image name")
	flag.Parse()
	cli = newCli(options)
}

func main() {
	cli.run(newImage("vim_dishwash_bar.jpg"))
}
