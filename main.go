package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"

	"os"
)

var (
	fishColor   = color.RGBA{0x40, 0xE0, 0xD0, 0xff}
	sharkColor  = color.RGBA{0x99, 0x32, 0xCC, 0xff}
	waterColor  = color.RGBA{0x00, 0x00, 0x80, 0xff}
	delay       = 6
	numChronons = 500
	outputFile  = "wator.gif"
)

func main() {

	var images []*image.Paletted
	var delays []int
	initWorld()
	img := colorWorld()
	images = append(images, img)
	delays = append(delays, delay)
	for i := 0; i < numChronons; i++ {
		nextChronon()
		img := colorWorld()
		images = append(images, img)
		delays = append(delays, 10)
	}

	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//gif.Encode(f, img, nil)
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

}
