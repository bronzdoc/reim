package image

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func clean() {
	currentDir, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer currentDir.Close()

	files, err := currentDir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".jpeg" {
			os.Remove(file.Name())
		}

	}
}

func TestNewImage(t *testing.T) {
	defer clean()
	newImage := NewImage("../../test/fixtures/jpeg_image.jpeg")
	imageName := "jpeg_image.jpeg"
	if newImage.Name() == imageName {
		t.Error("Expected %s to != %s", newImage.Name(), imageName)
	}

	newImageExtension := newImage.ImageFormat.Extension()
	if newImageExtension != ".jpeg" {
		t.Error("Expected %s to != %s", newImageExtension, ".jpeg")
	}

}

func TestRename(t *testing.T) {
	defer clean()
	newImage := NewImage("../../test/fixtures/jpeg_image.jpeg")
	newImage.Rename("new_image_name.jpeg")
	imageName := newImage.Name()
	if imageName != "new_image_name.jpeg" {
		t.Error("Expected %s to != %s", imageName, "new_image_name.jpeg")
	}
}

func TestResize(t *testing.T) {
	defer clean()
	newImage := NewImage("../../test/fixtures/jpeg_image.jpeg")
	newImage.Resize(133, 133)
	if newImage.Width != 133 && newImage.Height != 133 {
		t.Error("Expected image to have width and heigth == 133")
	}
}
