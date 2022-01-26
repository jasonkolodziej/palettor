package palettor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jasonkolodziej/palettor/hex"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// base64-encoded 4x4 png, w/ black, white, red, & blue pixels
var testImageData = []byte("iVBORw0KGgoAAAANSUhEUgAAAAIAAAACCAIAAAD91JpzAAAAE0lEQVQIHWMAgv///zP8ZwCC/wAh7AT8vKm73AAAAABJRU5ErkJggg==")

func TestExtract(t *testing.T) {
	decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(testImageData))
	img, err := png.Decode(decoder)
	if err != nil {
		t.Fatalf("invalid test image: %s", err)
	}

	_, err = Extract(5, 100, img)
	if err == nil {
		t.Errorf("k too large, expected an error")
	}

	palette, _ := Extract(4, 100, img)
	if palette.Count() != 4 {
		t.Errorf("expected 4 colors, got %d", palette.Count())
	}
}

func TestPalette_Colors(t *testing.T) {
	d, err := ioutil.ReadFile("testdata/test300x300_spotify-singles.jpeg")
	if err != nil {
		fmt.Println("File reading error", err)
		return
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
	palette, _ := Extract(5, 100, img)

	// Palette is a mapping from color to the weight of that color's cluster,
	// which can be used as an approximation for that color's relative
	// dominance
	for _, c := range palette.Colors() {
		fmt.Printf("color: %v; hexCode: %s; weight: %v\n", c, hex.ColorToHex(c), palette.Weight(c))
		if err := json.NewEncoder(os.Stdout).Encode(c); err != nil {
			log.Fatalf("Error encoding JSON: %s", err)
		}

	}
}
