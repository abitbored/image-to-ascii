package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

var asciiChars = []string{"@", "#", "S", "%", "?", "*", "+", ";", ":", ",", "."}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <image.jpg>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: could not open image file")
		return
	}

	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error: could not decode image")
		return
	}

	newWidth := 100
	aspectRatio := 0.1 // change this part to adjust the image so the image does not appear squished
	newHeight := uint(float64(img.Bounds().Dy()) * aspectRatio)
	img = resize.Resize(uint(newWidth), newHeight, img, resize.Lanczos3)

	asciiArt := imageToASCII(img)
	fmt.Println(asciiArt)
}

func imageToASCII(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var result string

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			asciiChar := grayToASCII(grayColor.Y)
			result += asciiChar
		}
		result += "\n"
	}

	return result
}

func grayToASCII(gray uint8) string {
	scale := float64(gray) / 255.0
	index := int(scale * float64(len(asciiChars)-1))
	return asciiChars[index]
}
