package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func openImage(path string) ([]*image.Paletted, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gifka, err := gif.DecodeAll(f)
	if err != nil {
		return nil, err
	}

	return gifka.Image, nil
}

func avgPixel(img *image.Paletted, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayscale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

func printASCII(img *image.Paletted) {
	ramp := "@#+=. "
	max := img.Bounds().Max
	scaleX, scaleY := 10, 5
	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := avgPixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*c/65536]))
		}
		fmt.Println()
	}
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func main() {
	path := flag.String("path", "", "path to gif file")
	flag.Parse()

	palete, err := openImage(*path)
	if err != nil {
		return
	}
	for _, img := range palete {
		printASCII(img)
	}
}
