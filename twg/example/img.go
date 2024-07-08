package example

import "io"

// Image is a struct that represents an image.
type Image struct{}

// Decode is a fake function that decodes an image from an io.Reader.
func Decode(r io.Reader) (*Image, error) {
	return &Image{}, nil
}

// Crop is a fake function that crops an image.
func Crop(img *Image, x1, y1, x2, y2 int) error {
	return nil
}

// Encode is a fake function that encodes an image to an io.Writer.
func Encode(img *Image, format string, w io.Writer) error {
	return nil
}
