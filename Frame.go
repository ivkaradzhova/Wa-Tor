package main

import (
	"math/rand"
)

var (
	world2      [length][width]byte
	starvation2 [length][width]byte
	breed2      [length][width]byte
	energy2     [length][width]byte
)

func killAnimal(x int, y int) {
	world[x][y] = water
	starvation[x][y] = 0
	breed[x][y] = 0
	energy[x][y] = 0
	world2[x][y] = water
	starvation2[x][y] = 0
	breed2[x][y] = 0
	energy2[x][y] = 0

}

func updateEnergy(x int, y int, newX int, newY int) bool {
	if energy[x][y]-1 == 0 {
		killAnimal(x, y)
		return true
	} else {
		energy2[newX][newY] = energy[x][y] - 1
		return false
	}
}

func moveAnimal(x int, y int, newX int, newY int, animal byte) {
	//updateEnergy(x,y, newX, newY)
	//if animal == fish {
	//	updateFishStarvation(x,y, newX, newY)
	//}
	world[x][y] = water
	world2[newX][newY] = animal
	//updateBreed(x , y, newX, newY, animal)
}

func updateBreed(x int, y int, newX int, newY int, animal byte) {
	if breed[x][y] <= 0 {
		//new fish/shark is created
		world2[x][y] = animal
		if animal == shark {
			breed2[x][y] = sharkBreed
			starvation2[x][y] = starvationTime
		} else {
			breed2[x][y] = fishBreed
			energy2[x][y] = energyTime
		}
		//update the parent fish/shark
		if animal == shark {
			breed2[newX][newY] = sharkBreed
		} else {
			breed2[newX][newY] = fishBreed
		}
	} else {
		breed2[newX][newY] = breed[x][y] - 1
		breed[x][y] = 0
	}
}

func getDirections(x int, y int) (int, int, int, int) {
	north := x - 1
	south := x + 1
	east := y + 1
	west := y - 1
	if north < 0 {
		north = length - 1
	}
	if south >= length {
		south = 0
	}
	if east >= width {
		east = 0
	}
	if west < 0 {
		west = width - 1
	}
	return north, south, east, west
}

func findAdjacentCells(x int, y int, target byte) [][2]int {
	north, south, east, west := getDirections(x, y)
	var freeCells [][2]int
	if world[north][y] == target {
		freeCells = append(freeCells, [2]int{north, y})
	}
	if world[x][east] == target {
		freeCells = append(freeCells, [2]int{x, east})
	}
	if world[south][y] == target {
		freeCells = append(freeCells, [2]int{south, y})
	}
	if world[x][west] == target {
		freeCells = append(freeCells, [2]int{x, west})
	}
	return freeCells
}

func updateFishStats(x int, y int, newX int, newY int) {
	died := updateEnergy(x, y, newX, newY)
	if !died {
		moveAnimal(x, y, newX, newY, fish)
		updateBreed(x, y, newX, newY, fish)
	}
}

func handleFish(x int, y int) {
	freeCells := findAdjacentCells(x, y, water)
	if freeCells == nil {
		updateFishStats(x, y, x, y)

	} else {
		direction := freeCells[rand.Intn(len(freeCells))]
		updateFishStats(x, y, direction[0], direction[1])
	}
}

func updateStarvation(starve bool, x int, y int, newX int, newY int) bool {
	if starve {
		if starvation[x][y]-1 == 0 {
			killAnimal(x, y)
			return true
		} else {
			starvation2[newX][newY] = starvation[x][y] - 1
			starvation[x][y] = 0
		}
	} else {
		starvation2[newX][newY] = starvation[x][y] + 1
		starvation[x][y] = 0
	}
	return false
}

func updateSharkStats(x int, y int, newX int, newY int, isEating bool) {
	if died := updateStarvation(isEating, x, y, newX, newY); !died {
		updateBreed(x, y, newX, newY, shark)
		moveAnimal(x, y, newX, newY, shark)
	}
}

func handleShark(x int, y int) {
	fishCells := findAdjacentCells(x, y, fish)
	if fishCells != nil {
		directions := fishCells[rand.Intn(len(fishCells))]
		killAnimal(directions[0], directions[1])
		updateSharkStats(x, y, directions[0], directions[1], true)
	} else {
		freeCells := findAdjacentCells(x, y, water)
		if freeCells != nil {
			directions := freeCells[rand.Intn(len(freeCells))]
			updateSharkStats(x, y, directions[0], directions[1], false)
		} else {
			updateSharkStats(x, y, x, y, false)
		}
	}

}

func nextChronon() {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			switch world[i][j] {
			case fish:
				handleFish(i, j)
			case shark:
				handleShark(i, j)
			}
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			world[i][j] = world2[i][j]
			starvation[i][j] = starvation2[i][j]
			breed[i][j] = breed2[i][j]
			energy[i][j] = energy2[i][j]
		}
	}
	//fillWithWater(world2)
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			world2[i][j] = water
			starvation2[i][j] = 0
			breed2[i][j] = 0
			energy2[i][j] = 0
		}
	}
}
