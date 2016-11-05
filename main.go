package main

import (
	"flag"
	"github.com/bronzdoc/slacky/lib"
	"github.com/bronzdoc/slacky/lib/image"
)

type Options struct {
	Width     uint
	Height    uint
	Out       string
	ImageName string
}

type Cli struct {
	Options Options
}

func newCli(options Options) *Cli {
	return &Cli{Options: options}
}

func (c *Cli) run(image *image.Image) {
	width := c.Options.Width
	height := c.Options.Height
	image.Resize(width, height, "slack-image.jpg")
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
	formater := new(lib.JPEG)
	cli.run(
		image.NewImage(
			"/home/bronzdoc/Pictures/vim_dishwash_bar.jpg",
			formater,
		),
	)
}
