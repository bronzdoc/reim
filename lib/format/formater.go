package format

import (
	"image"
	"io"
	"os"
)

type Formater interface {
	Encode(io.Writer, image.Image)
	Decode(*os.File) (image.Image, error)
	MagicNumber() string
}
