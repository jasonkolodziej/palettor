package palettor

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
)

func Grab(fileName string, kDominantColors, maxIterations int) *Palette {
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	img, err := jpeg.Decode(bytes.NewReader(d))
	if err != nil {
		log.Fatal(err)
	}

	// For a real-world use case, it's best to use something like
	// github.com/nfnt/resize to transform images into a manageable size before
	// extracting colors:
	//
	//     img = resize.Thumbnail(200, 200, img, resize.Lanczos3)
	//
	// In this example, we're already starting from a tiny image.

	// Extract the 3 most dominant colors, halting the clustering algorithm
	// after 100 iterations if the clusters have not yet converged.
	palette, err := Extract(kDominantColors, maxIterations, img)

	if err != nil {
		log.Fatal(err)
	}

	return palette
}
