package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	// "log/slog"
)

var debug bool = true
var score int = 0

const MAX_X = 4
const MAX_Y = 4

const blockSize = 128

// List of All States
const (
	IDLE = iota
	START

	// MOVEMENTS STATES
	MOVE_RIGHT
	MOVE_UP
	MOVE_DOWN
	MOVE_LEFT

	MOVE_ZERO_RIGHT
	MOVE_ZERO_LEFT
	MOVE_ZERO_DOWN
	MOVE_ZERO_UP

	JUST_FINISH
	LOSE
)

var tileMap [MAX_X][MAX_Y]int = [MAX_X][MAX_Y]int{
	{0, 0, 0, 0},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
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
				rl.DrawRectangle(int32(i3*(blockSize+1)), int32(i*(blockSize+1)), blockSize, blockSize, rl.Gray)
			} else {
				rl.DrawRectangle(int32(i3*(blockSize+1)), int32(i*(blockSize+1)), blockSize, blockSize, rl.ColorFromNormalized(rl.Vector4{float32(value) * 5, float32(value) * 15, float32(value) * 3, float32(value)}))
				rl.DrawText(fmt.Sprintf("%d", value), int32(i3)*blockSize+(blockSize/2)-5, int32(i)*blockSize+(blockSize/2)-15, 30, rl.Black)
			}
		}
	}
}

// func moveItems() {
// 	var all_clean bool = true
// 	for k, v := range State.GetDataFrame() {
// 		if v == 1 {
// 			move(k)
// 			all_clean = false
// 		}
// 	}
//
// 	// State.SetState(IDLE)
// }

func process_state() int {

	switch State.GetState() {
	// case START:
	// case IDLE:
	case MOVE_RIGHT:
		// slog.Info("ASD")
		Move_v(0, +1, MAX_X-1)
		State.SetState(MOVE_ZERO_RIGHT)
	case MOVE_LEFT:
		Move_v(3, -1, 0)
		State.SetState(MOVE_ZERO_LEFT)
	case MOVE_UP:
		Move_h(0, +1, MAX_Y-1)
		State.SetState(MOVE_ZERO_UP)
	case MOVE_DOWN:
		Move_h(3, -1, 0)
		State.SetState(MOVE_ZERO_DOWN)
		// move_zero_v(0, +1, MAX_X-1)

	case MOVE_ZERO_RIGHT:
		isClean := Move_zero_v(0, +1, MAX_X-1)
		if isClean {
			State.SetState(JUST_FINISH)
		}

	case MOVE_ZERO_LEFT:
		isClean := Move_zero_v(3, -1, 0)
		if isClean {
			State.SetState(JUST_FINISH)
		}

	case MOVE_ZERO_DOWN:
		isClean := Move_zero_h(3, -1, 0)
		if isClean {
			State.SetState(JUST_FINISH)
		}

	case MOVE_ZERO_UP:
		isClean := Move_zero_h(0, +1, MAX_Y-1)
		if isClean {
			State.SetState(JUST_FINISH)
		}

	case JUST_FINISH:
		Add_item()
		State.SetState(IDLE)
		// add_item()
		// State.SetState(IDLE)
	case IDLE:
		if rl.IsKeyPressed(int32(Get_key(RIGHT))) {
			State.SetState(MOVE_RIGHT)
		} else if rl.IsKeyPressed(int32(Get_key(LEFT))) {
			// leftAvail = true
			State.SetState(MOVE_LEFT)
		} else if rl.IsKeyPressed(int32(Get_key(UP))) {
			// upAvail = true
			State.SetState(MOVE_UP)
		} else if rl.IsKeyPressed(int32(Get_key(BOTTOM))) {
			// bottomAvail = true
			State.SetState(MOVE_DOWN)
		} else if rl.IsKeyPressed(rl.KeyEscape) {
			return 2
		}
		// if rl.IsKeyPressed(rl.KeyB) {
		// 	Add_item()
		// }
	}
	// Do things based on the current State
	return 0
}

func setupGame() {
	// TODO: This will be a menu to start the game
	// But for now we will start the game immediatelly
	State.SetState(IDLE)
}

func main() {
	rl.InitWindow(515, 480+30, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	setupGame()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		drawMap()

		ret_process := process_state()

		// Exit the game
		if ret_process == 2 {
			break
		}

		rl.EndDrawing()
	}
}
