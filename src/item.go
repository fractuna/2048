package main

import (
	"fmt"
	"math/rand/v2"
)

// TODO: it needs more work
func Add_item() {
	var candid_columns []Tuple

	// Get all the coulmns that has empty room inside it
	x, y := 0, 0
	for {
		if x == 3 {
			break
		}

		if y == 4 {
			x++
			y = 0
		}

		if tileMap[x][y] == 0 {
			fmt.Printf("DEBUG: Found candidate for add_item x: %d, y: %d\n", x, y)
			candid_columns = append(candid_columns, newTuple(x, y))
		}
		y++
	}

	// This means nohting is there for adding a new item
	if len(candid_columns) == 0 {
		return
	}

	fmt.Printf("DEBUG: The length of the candidate_items array is: %d\n", len(candid_columns))
	random_horizantal := candid_columns[(rand.IntN((len(candid_columns)))+1)-1]
	fmt.Printf("DEBUG: choosed the %d from candidates column\n", random_horizantal)

	tileMap[random_horizantal.getFirst()][random_horizantal.getSecond()] = 1
	l_item += 1
	fmt.Println("I added a new item")
}

/* A simple function that checks if player can move an item or not */
func possible_move() bool {
	move_right := Move_v(0, +1, MAX_X-1)

	move_left := Move_v(3, -1, 0)
	move_douw := Move_h(0, +1, MAX_Y-1)

	move_up := Move_h(3, -1, 0)

	return (move_right && move_left && move_douw && move_up)
}

/* This function works by shifting items to fill the zero values*/
func Move_zero_v(default_x int, dmp int, max_x int) bool {
	x := default_x
	y := 0
	zeroClean := true
	for {
		if y == 4 {
			return zeroClean
		}
		x = default_x
		// x = default_x
		for {
			if x == max_x {
				break
			}
			if (tileMap[y][x+dmp] == 0) && tileMap[y][x] != 0 {
				tileMap[y][x+dmp] = tileMap[y][x]
				tileMap[y][x] = 0
				// fmt.Println(tileMap[y])
				zeroClean = false
				// y += 1
				break
			}

			x += dmp
			// x += dmp
		}
		y += 1
	}
}

/* This function works by shifting items to fill the zero values*/
func Move_zero_h(default_y int, dmp int, max_y int) bool {
	x := 0
	y := default_y
	zeroClean := true
	for {
		if x == 4 {
			return zeroClean
		}
		y = default_y
		for {
			if y == max_y {
				break
			}
			if (tileMap[y+dmp][x] == 0) && tileMap[y][x] != 0 {
				tileMap[y+dmp][x] = tileMap[y][x]
				tileMap[y][x] = 0
				// fmt.Println(tileMap[y])
				zeroClean = false
				break
			}
			y += dmp
		}
		x += 1
	}
}

// Original move right function
func Move_v(default_x int, dmp int, max_x int) bool {
	x := default_x
	y := 0
	var isClean bool = true
	for {
		if y == 4 {
			return isClean
		}
		x = default_x
		for {
			if x == max_x {
				break
			}
			if (tileMap[y][x] == tileMap[y][x+dmp]) && (tileMap[y][x] != 0) {

				fmt.Println(tileMap[y][x], tileMap[y][x+dmp], y, x, x+dmp)
				// INFO: Just for specefic state
				if State.GetState() == CEHCK_ITEMS {
					return false // == (isClean = false)
				}

				tileMap[y][x+dmp] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				l_item -= 1
				isClean = false
				break
			}
			x += dmp
		}
		y += 1
	}
}

func Move_h(default_y int, dmp int, max_y int) bool {
	x := 0
	y := default_y
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

				// INFO: Just for specefic state
				if State.GetState() == CEHCK_ITEMS {
					return false
				}
				// if tileMap[y][x+1] == 0 && tileMap[y][x] != 0 {
				tileMap[y+dmp][x] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				l_item -= 1
				// x = 0
				// fmt.Println(tileMap[y])
				isClean = false
				x += 1
				break
			}
			y += dmp
		}
	}
}
