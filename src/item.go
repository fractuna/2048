package main

import (
	"math/rand/v2"
)

// TODO: it needs more work
func Add_item() {
	// make a new item
	// find a place for the item
	// TODO: It's better to choose which column has chance
	// To be picked rather than using randomize numbers
	var random_horizantal = rand.IntN((MAX_Y-1)-0) + 0
	// var random_horizantal = 1
	// fmt.Println("That's the random colmn:", random_horizantal)

	// First check if the colmn has empty rooms

	// var dmp int = 0
	for i := MAX_Y - 1; i >= 0; i-- {
		item := tileMap[i][random_horizantal]
		// fmt.Println(tileMap[i][random_horizantal])
		if item == 0 {
			// TODO: Add add_item logic here...
			// backward until find an item
			tileMap[i][random_horizantal] = 1
			// tileMap[i+dmp][random_horizantal] = 1
			break
		}
		// // maybe it doesn't have any item
		// if i == 0 {
		// 	tileMap[MAX_Y-1][random_horizantal] = 1
		// 	break
		// }
	}
}

/* This function works by shifting items to fill the zero values*/
func Move_zero_v(default_x int, dmp int, max_x int) bool {
	var x = default_x
	var y = 0
	zeroClean := true
	for {
		if y == 4 {
			return zeroClean
		}
		for {
			if x == max_x {
				x = default_x
				y += 1
				break
			}
			if (tileMap[y][x+dmp] == 0) && tileMap[y][x] != 0 {
				tileMap[y][x+dmp] = tileMap[y][x]
				tileMap[y][x] = 0
				// fmt.Println(tileMap[y])
				zeroClean = false
				y += 1
				break
			}
			x += dmp
		}
	}
}

/* This function works by shifting items to fill the zero values*/
func Move_zero_h(default_y int, dmp int, max_y int) bool {
	var x = 0
	var y = default_y
	zeroClean := true
	for {
		if x == 4 {
			return zeroClean
		}
		for {
			if y == max_y {
				y = default_y
				x += 1
				break
			}
			if (tileMap[y+dmp][x] == 0) && tileMap[y][x] != 0 {
				tileMap[y+dmp][x] = tileMap[y][x]
				tileMap[y][x] = 0
				// fmt.Println(tileMap[y])
				zeroClean = false
				x += 1
				break
			}
			y += dmp
		}
	}
}

// Original move right function
func Move_v(default_x int, dmp int, max_x int) bool {
	var x = 0
	var y = 0
	var isClean bool = true
	for {
		if y == 4 {
			// if isClean {
			// 	// TODO: Maybe this written better ;d: got it :D
			// 	// rightAvail = false
			// }
			return isClean
		}
		x = default_x
		for {
			if x == max_x {
				// x = default_x
				y += 1
				break
			}
			if (tileMap[y][x] == tileMap[y][x+dmp]) && (tileMap[y][x] != 0) {
				// if tileMap[y][x+1] == 0 && tileMap[y][x] != 0 {
				tileMap[y][x+dmp] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				// x = 0
				// fmt.Println(tileMap[y])
				y += 1
				break
			}
			// } else if (tileMap[y][x+dmp] == 0) && tileMap[y][x] != 0 {
			// 	tileMap[y][x+dmp] = tileMap[y][x]
			// 	tileMap[y][x] = 0
			// 	// fmt.Println(tileMap[y])
			// 	y += 1
			// 	isClean = false
			// 	break
			// } else {
			x += dmp
			// }
		}
	}
}

func Move_h(default_y int, dmp int, max_y int) bool {
	var x = 0
	var y = 0
	var isClean bool = true
	for {
		if x == 4 {
			// if isClean {
			// 	// TODO: Maybe this written better ;d: got it :D
			// 	// rightAvail = false
			// }
			return isClean
		}
		y = default_y
		for {
			if y == max_y {
				// y = default_y
				x += 1
				break
			}
			if (tileMap[y][x] == tileMap[y+dmp][x]) && (tileMap[y][x] != 0) {
				// if tileMap[y][x+1] == 0 && tileMap[y][x] != 0 {
				tileMap[y+dmp][x] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				// x = 0
				// fmt.Println(tileMap[y])
				x += 1
				break
			}
			y += dmp
		}
	}
}
