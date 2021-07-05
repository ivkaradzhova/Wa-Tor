package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

const (
	fish  = 1
	shark = 2
	water = 0
)
const (
	length = 490
	width  = 980
)

const (
	fishBreed      = 10
	sharkBreed     = 9
	starvationTime = 11
	energyTime     = 20
	numFish        = 1000
	numSharks      = 100
)

var (
	world      [length][width]byte
	starvation [length][width]byte
	energy     [length][width]byte
	breed      [length][width]byte
)

func generateRand(min int, max int, n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(max-min) + min
	}
	return res
}

func addAnimal(num int, animal byte) {
	xCord := generateRand(0, length, num)
	yCord := generateRand(0, width, num)

	if animal == fish {
		for i := 0; i < num; i++ {
			world[xCord[i]][yCord[i]] = fish
			breed[xCord[i]][yCord[i]] = fishBreed
			energy[xCord[i]][yCord[i]] = energyTime
		}
	} else if animal == shark {
		for i := 0; i < num; i++ {
			world[xCord[i]][yCord[i]] = shark
			starvation[xCord[i]][yCord[i]] = starvationTime
			breed[xCord[i]][yCord[i]] = sharkBreed
		}
	}

}

func fillWithWater(world [length][width]byte) {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			world[i][j] = water
		}
	}
}

func initWorld() {
	rand.Seed(time.Now().UnixNano())
	fillWithWater(world)
	fillWithWater(world2)
	addAnimal(numFish, fish)
	addAnimal(numSharks, shark)
}

func colorWorld() *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, width*2, length*2), color.Palette{waterColor, fishColor, sharkColor})
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			switch world[i][j] {
			case fish:
				img.Set(j*2, i*2, fishColor)
				img.Set(j*2+1, i*2, fishColor)
				img.Set(j*2, i*2+1, fishColor)
				img.Set(j*2+1, i*2+1, fishColor)

			case shark:
				img.Set(j*2, i*2, sharkColor)
				img.Set(j*2+1, i*2, sharkColor)
				img.Set(j*2, i*2+1, sharkColor)
				img.Set(j*2+1, i*2+1, sharkColor)
			}
		}
	}
	return img
}
