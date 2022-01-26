// Shortest GUI program written in Golang.
// It displays a window and exits when the "close" button of the window is clicked.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jasonkolodziej/palettor"
	"github.com/jasonkolodziej/palettor/hex"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
)

const fileName = "testdata/test300x300_spotify-singles.jpeg"

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gopher Tests")

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Aurélie Vache"),
			), myWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)

	// Define a welcome text centered
	text := canvas.NewText("Origin Photo", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, _ = fyne.LoadResourceFromPath(fileName)
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 300, Height: 300}) // by default size is 0, 0
	gopherImg.FillMode = canvas.ImageFillOriginal
	var p1, p2 = Grab(fileName, 10, 300)

	// Define a "random" button
	//randomBtn := widget.NewButton("Run", func() {
	//	//resource, _ := fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
	//	//gopherImg.Resource = resource
	//	p1, p2 = Grab(fileName, 5, 100)
	//
	//	//Redrawn the image with the new path
	//	// gopherImg.Refresh()
	//})
	//randomBtn.Importance = widget.HighImportance

	// Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		gopherImg,
		p1,
		p2,
		//randomBtn,
	)

	// Display our content
	myWindow.SetContent(box)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}

func Grab(fileName string, kDominantColors, maxIterations int) (*widget.Card, *widget.Card) {
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, nil
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
	palette, err := palettor.Extract(kDominantColors, maxIterations, img)

	if err != nil {
		log.Fatal(err)
	}
	// Palette is a mapping from color to the weight of that color's cluster,
	// which can be used as an approximation for that color's relative
	// dominance
	//
	ogBox := container.NewHBox()
	box := container.NewHBox()
	l := len(palette.Colors())
	fmt.Printf("number of Colors is : %d; width set to: %f\n", l, float32(300/l))
	var size = fyne.Size{
		Width:  float32(300/l) / 2,
		Height: 25,
	}
	var size1 = fyne.Size{
		Width:  float32(300/l) / 2,
		Height: 25,
	}
	for _, c := range palette.Colors() {
		r := canvas.NewRectangle(c)
		r.SetMinSize(size)
		c1 := hex.FromColor(c).ToRGBA()
		r1 := canvas.NewRectangle(c1)
		r1.SetMinSize(size1)
		ss, _ := json.Marshal(c)
		ss1, _ := json.Marshal(c1)
		ogBox.Add(widget.NewCard(fmt.Sprintf("weight: %4f", palette.Weight(c)), fmt.Sprintf("color: %s", ss), r))
		box.Add(widget.NewCard("", fmt.Sprintf("color: %s", ss1), r1))
	}

	return widget.NewCard("Original", "", ogBox), widget.NewCard("RGB", "", box)
}

//import "github.com/muesli/gamut"
//import "github.com/muesli/gamut/palette"
//import "github.com/muesli/gamut/theme"
//import "github.com/mandykoh/prism"
