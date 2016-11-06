package lib

import (
	"github.com/bronzdoc/slacky/lib/image"
)

type Options struct {
	Width     uint
	Height    uint
	Out       string
	ImageName string
}

type cli struct {
	Options Options
}

func NewCli(options Options) *cli {
	return &cli{Options: options}
}

func (c *cli) Run() {
	width := c.Options.Width
	height := c.Options.Height
	name := c.Options.ImageName
	newName := c.Options.Out

	image := image.NewImage(name)

	if newName == "" {
		newName = "slack-image.jpg"
	}

	image.Resize(width, height, newName)
}
