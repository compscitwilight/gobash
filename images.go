package main

import (
	"image"
	"os"
)

func decodeImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	return image, err
}
