package example_test

import (
	"fmt"
	"io"

	// Needed for initialize side effect
	_ "image/png"

	"github.com/piyushbag/twg/example"
)

var file string = "this is not used."

func Example_crop() {
	var r io.Reader
	img, err := example.Decode(r)
	if err != nil {
		panic(err)
	}

	if err := example.Crop(img, 0, 0, 100, 100); err != nil {
		panic(err)
	}

	var w io.Writer
	if err := example.Encode(img, "jpeg", w); err != nil {
		panic(err)
	}
	fmt.Println("See out.jpg for the cropped image.")

	// Output:
	// See out.jpg for the cropped image.
}
