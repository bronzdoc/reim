package main

import (
	"flag"
	"fmt"
	"github.com/bronzdoc/reim/cli"
	"os"
)

var (
	options cli.Options
)

func init() {
	flag.Usage = func() {
		fmt.Printf("%s - image resizer\n", os.Args[0])
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    reim [options] file...\n\n")
		flag.PrintDefaults()
	}

	options = cli.Options{}
	flag.UintVar(&options.Width, "width", 128, "New image width")
	flag.UintVar(&options.Height, "height", 128, "New image height")
	flag.StringVar(&options.Out, "out", "", "Generated image name")

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
