package main

import (
	"fmt"
	// "strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const maxX = 129
const maxY = 129

var rightAvail bool = false

const blockSize = 128

// const blockPadding = 8

var tileMap [4][4]int = [4][4]int{
	{1, 1, 0, 2},
	{1, 1, 1, 1},
	{1, 1, 0, 1},
	{1, 0, 1, 1},
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
			// fmt.Println(i, fmt.Sprintf("%d, %d", i, int32(i)))
			rl.DrawRectangle(int32(i3*maxX), int32(i*maxY), blockSize, blockSize, rl.Blue)
			rl.DrawText(fmt.Sprintf("%d", value), int32(i3)*blockSize+(blockSize/2)-5, int32(i)*blockSize+(blockSize/2)-15, 30, rl.Black)
			// rl.DrawText(fmt.Sprintf("%d", value), int32(i3)*blockSize, int32(i)*blockSize, 30, rl.Black)
		}
	}
}

func moveItems() {
	if rightAvail {
		_move_right()
	}
}

// Original move right function
func _move_right() {
	var x = 0
	var y = 0
	for {
		if y == 4 {
			return
		}
		for {
			if x == 3 {
				x = 0
				y += 1
				break
			}
			if (tileMap[y][x] == tileMap[y][x+1]) && (tileMap[y][x] != 0) {
				// if tileMap[y][x+1] == 0 && tileMap[y][x] != 0 {
				tileMap[y][x+1] = tileMap[y][x] * 2
				tileMap[y][x] = 0
				x = 0
				fmt.Println(tileMap[y])
				y += 1
				break
			} else if (tileMap[y][x+1] == 0) && tileMap[y][x] != 0 {
				tileMap[y][x+1] = tileMap[y][x]
				tileMap[y][x] = 0
				x = 0
				fmt.Println(tileMap[y])
				y += 1
				break
			} else {
				x += 1
			}
		}
	}
}

func main() {
	rl.InitWindow(515, 480+30, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	// createMap()
	// moveRight()
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		drawMap()

		if rl.IsKeyPressed(rl.KeyRight) {
			rightAvail = true
		}

		moveItems()

		rl.EndDrawing()
	}
}
