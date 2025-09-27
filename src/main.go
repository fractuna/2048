package main

import (
	"fmt"
	"math/rand/v2"
	// "strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Simple enum for movements structure
const (
	RIGHT = iota
	LEFT
	UP
	BOTTOM
)

const (
	IDLE = iota
	START
	MOVING
	JUST_FINISH
	WIN
	LOSE
)

const MAX_X = 4
const MAX_Y = 4
const maxX = 129
const maxY = 129

var upAvail bool = false
var rightAvail bool = false
var leftAvail bool = false
var bottomAvail bool = false

const blockSize = 128

// const blockPadding = 8

var tileMap [MAX_X][MAX_Y]int = [MAX_X][MAX_Y]int{
	{0, 0, 0, 0},
	{1, 1, 0, 0},
	{1, 1, 0, 0},
	{1, 1, 1, 1},
}

func createMap() {
	tileMap[0][0] = 1
	tileMap[0][1] = 0
	tileMap[0][2] = 1
	tileMap[0][3] = 5
}

func drawMap() {
	for i, v := range tileMap {
		for i3, value := range v {
			if value == 0 {
				rl.DrawRectangle(int32(i3*maxX), int32(i*maxY), blockSize, blockSize, rl.Gray)
			} else {
				rl.DrawRectangle(int32(i3*maxX), int32(i*maxY), blockSize, blockSize, rl.Beige)
				rl.DrawText(fmt.Sprintf("%d", value), int32(i3)*blockSize+(blockSize/2)-5, int32(i)*blockSize+(blockSize/2)-15, 30, rl.Black)
			}
		}
	}
}

func moveItems() {
	for k, v := range State.getDataFrame() {
		if v == 1 {
			move(k)
		}
	}
	// switch State.getData(
	// if rightAvail {
	// 	move(RIGHT)
	// 	return
	// }
	// if leftAvail {
	// 	move(LEFT)
	// 	return
	// }
	// if upAvail {
	// 	move(UP)
	// 	return
	// }
	// if bottomAvail {
	// 	move(BOTTOM)
	// 	return
	// }
	State.setState(IDLE)
}

func item_length() (count int) {
	for i, items := range tileMap {
		for i3, _ := range items {
			if tileMap[i][i3] != 0 {
				count += 1
			}
		}
	}
	return count
}

// TODO: it needs more work
func add_item() {
	// make a new item
	// find a place for the item
	var random_horizantal = rand.IntN(MAX_Y - 1)
	fmt.Println("That's the random colmn:", random_horizantal)
	for i := range MAX_Y {
		item := tileMap[i][random_horizantal]
		if item != 0 {
			fmt.Println("Added a new item")
			tileMap[i-1][random_horizantal] = 1
			break
		}
		if i == MAX_Y-1 {
			tileMap[i][random_horizantal] = 1
			break
		}
	}
}

// TODO: Maybe we can merge v,h functions together :d
func move(direction int) {
	// default values for all directions
	var default_mov_pointer_x = 0
	var default_mov_pointer_y = 0
	var default_x = 0
	var default_y = 0
	var isClean bool = false

	switch direction {
	case RIGHT:
		default_mov_pointer_x = +1
		default_x = 0
		isClean = _move_v(default_x, default_mov_pointer_x, MAX_X-1)
		State.setData(RIGHT, Bool2int(!isClean))
	case LEFT:
		default_mov_pointer_x = -1
		default_x = MAX_X - 1
		isClean = _move_v(default_x, default_mov_pointer_x, 0)
		State.setData(LEFT, Bool2int(!isClean))
	case UP:
		default_mov_pointer_y = +1
		default_y = 0
		isClean = _move_h(default_y, default_mov_pointer_y, MAX_Y-1)
		State.setData(UP, Bool2int(!isClean))
	case BOTTOM:
		default_mov_pointer_y = -1
		default_y = MAX_Y - 1
		State.setData(BOTTOM, Bool2int(!_move_h(default_y, default_mov_pointer_y, 0)))
	}
}

// Original move right function
func _move_v(default_x int, dmp int, max_x int) bool {
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
				x = default_x
				y += 1
				break
			}
			if (tileMap[y][x] == tileMap[y][x+dmp]) && (tileMap[y][x] != 0) {
				// if tileMap[y][x+1] == 0 && tileMap[y][x] != 0 {
				tileMap[y][x+dmp] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				// x = 0
				fmt.Println(tileMap[y])
				y += 1
				isClean = false
				break
			} else if (tileMap[y][x+dmp] == 0) && tileMap[y][x] != 0 {
				tileMap[y][x+dmp] = tileMap[y][x]
				tileMap[y][x] = 0
				fmt.Println(tileMap[y])
				y += 1
				isClean = false
				break
			} else {
				x += dmp
			}
		}
	}
}

func process_State() int {

	switch State.getState() {
	// case START:
	// case IDLE:
	// case MOVING:
	case JUST_FINISH:
		add_item()
		State.setState(IDLE)

	}
	// Do things based on the current State
	return 0
}

func _move_h(default_y int, dmp int, max_y int) bool {
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
				y = default_y
				x += 1
				break
			}
			if (tileMap[y][x] == tileMap[y+dmp][x]) && (tileMap[y][x] != 0) {
				// if tileMap[y][x+1] == 0 && tileMap[y][x] != 0 {
				tileMap[y+dmp][x] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				// x = 0
				fmt.Println(tileMap[y])
				x += 1
				isClean = false
				break
			} else if (tileMap[y+dmp][x] == 0) && tileMap[y][x] != 0 {
				tileMap[y+dmp][x] = tileMap[y][x]
				tileMap[y][x] = 0
				fmt.Println(tileMap[y])
				x += 1
				isClean = false
				break
			} else {
				y += dmp
			}
		}
	}
}

func setupGame() {
	var mov_map_data map[int]int = map[int]int{
		RIGHT:  0,
		LEFT:   0,
		UP:     0,
		BOTTOM: 0,
	}
	State.setDataFrame(mov_map_data)
}

func main() {
	rl.InitWindow(515, 480+30, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	// createMap()
	// moveRight()
	setupGame()
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		drawMap()

		ret_process := process_State()
		if ret_process == 1 {
			continue
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			State.setData(RIGHT, 1)
			State.setState(MOVING)
		} else if rl.IsKeyPressed(rl.KeyLeft) {
			// leftAvail = true
			State.setData(LEFT, 1)
			State.setState(MOVING)
		} else if rl.IsKeyPressed(rl.KeyDown) {
			// upAvail = true
			State.setData(UP, 1)
			State.setState(MOVING)
		} else if rl.IsKeyPressed(rl.KeyUp) {
			// bottomAvail = true
			State.setData(BOTTOM, 1)
			State.setState(MOVING)
		}

		if rl.IsKeyPressed(rl.KeyB) {
			add_item()
		}

		// if item_length() < 4 {
		// 	add_item()
		// }

		moveItems()

		rl.EndDrawing()
	}
}
