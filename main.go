package main

import (
	"flag"
	"fmt"
	"github.com/bronzdoc/reim/cli"
	"os"
	"time"
)

var (
	options cli.Options
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    gozer [FLAGS] IMAGE_NAME\n\n")
		flag.PrintDefaults()
	}

	options = cli.Options{}
	flag.UintVar(&options.Width, "width", 20, "New image width")
	flag.UintVar(&options.Height, "height", 20, "New image height")

	defaultImageName := fmt.Sprint(time.Now().Unix())
	flag.StringVar(&options.Out, "out", defaultImageName, "Generated image name")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		options.ImageName = args[0]
	} else {
		flag.Usage()
	}
}

func main() {
	cli := cli.NewCli(options)
	cli.Run()
}
