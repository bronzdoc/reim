package main

import (
	"flag"
	"github.com/bronzdoc/slacky/cli"
)

var (
	options cli.Options
)

func init() {
	options = cli.Options{}
	flag.StringVar(&options.ImageName, "image", "", "Name of the image to be resized")
	flag.UintVar(&options.Width, "width", 20, "New image width")
	flag.UintVar(&options.Height, "height", 20, "New image height")
	flag.StringVar(&options.Out, "out", "slack-image.jpg", "new image name")
	flag.Parse()
}

func main() {
	cli := cli.NewCli(options)
	cli.Run()
}
