package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	// "log/slog"
)

var (
	score         int  = 0
	l_item        int  = 16 // based on the default game's map
	items_checked bool = false
	font_family   rl.Font
)

// === States ===
const (
	IDLE = iota
	START

	// MOVEMENTS STATES
	MOVE_RIGHT
	MOVE_UP
	MOVE_DOWN
	MOVE_LEFT

	CEHCK_ITEMS

	MOVE_ZERO_RIGHT
	MOVE_ZERO_LEFT
	MOVE_ZERO_DOWN
	MOVE_ZERO_UP

	JUST_FINISH
	LOSE
)

var tileMap [MAX_X][MAX_Y]int = [MAX_X][MAX_Y]int{
	{2, 1, 2, 1},
	{1, 2, 1, 2},
	{2, 1, 2, 1},
	{1, 4, 2, 1},
}

func add_score(isClean bool) {
	// It means if there *were* any
	// item merges then add the score
	if isClean == false {
		score += DEFAULT_SCORE
	}

	items_checked = false
}

func drawMap() {
	for i, v := range tileMap {
		for i3, value := range v {
			if value == 0 {
				rl.DrawRectangle(
					int32(i3*(BLOCK_SIZE+1)),
					int32(i*(BLOCK_SIZE+1)),
					BLOCK_SIZE,
					BLOCK_SIZE,
					rl.Gray,
				)
			} else {
				rl.DrawRectangle(int32(i3*(BLOCK_SIZE+1)), int32(i*(BLOCK_SIZE+1)), BLOCK_SIZE, BLOCK_SIZE, rl.ColorFromNormalized(rl.Vector4{float32(value) * 5, float32(value) * 5, float32(value) * 5, float32(value)}))
				x, y := calc_text_center(fmt.Sprintf("%d", value), BLOCK_SIZE, BLOCK_SIZE, H3)
				rl.DrawTextEx(font_family, fmt.Sprintf("%d", value), rl.NewVector2(float32(i3*BLOCK_SIZE+(x)), float32(i*BLOCK_SIZE+(y))), H3, 2, rl.Black)
			}
		}
	}
}

func render_score() {
	rl.DrawTextEx(font_family, fmt.Sprintf("%s %d", SCORE_TEXT, score), rl.NewVector2(5.0, 5.0), H4, 2, rl.Black)
}

// The first screen before start playing the game
// TODO: It needs some work
func start_screen_info() {
	x, y := calc_text_center(START_PLAY_TEXT, rl.GetScreenWidth(), rl.GetScreenHeight(), H4)
	rl.DrawTextEx(font_family, START_PLAY_TEXT, rl.NewVector2(float32(x), float32(y)), H4, 2, rl.Black)
}

func game_over() {
	x, y := calc_text_center(GAME_OVER_TEXT, rl.GetScreenWidth(), rl.GetScreenHeight(), H1)
	rl.DrawTextEx(font_family, GAME_OVER_TEXT, rl.NewVector2(float32(x), float32(y)), H1, 4, rl.Black)
	// rl.DrawText(GAME_OVER_TEXT, (int32(rl.GetScreenWidth()/2))-(rl.MeasureText(GAME_OVER_TEXT, DEFAULT_FONT_SIZE)), 510/2, DEFAULT_FONT_SIZE, rl.Black)
}

func process_state() int {
	var isClean bool = false
	switch State.GetState() {
	case START:
		start_screen_info()
		if rl.IsKeyPressed(int32(Get_key(PLAY))) {
			State.SetState(IDLE)
		}
	case MOVE_RIGHT:
		isClean = Move_v(0, +1, MAX_X-1)
		add_score(isClean)
		State.SetState(MOVE_ZERO_RIGHT)
	case MOVE_LEFT:
		isClean = Move_v(3, -1, 0)
		add_score(isClean)
		State.SetState(MOVE_ZERO_LEFT)
	case MOVE_UP:
		isClean = Move_h(0, +1, MAX_Y-1)
		add_score(isClean)
		State.SetState(MOVE_ZERO_UP)
	case MOVE_DOWN:
		isClean = Move_h(3, -1, 0)
		add_score(isClean)
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
		if !isClean {
			Add_item()
		}
		State.SetState(IDLE)
	case LOSE:
		game_over()
		// return 2
	case IDLE:

		if l_item >= L_MAX {
			if items_checked == false {
				// Check if the player can't do any movement
				State.SetState(CEHCK_ITEMS)
				if possible_move() {
					State.SetState(LOSE)
				} else {
					State.SetState(IDLE)
				}
				items_checked = true
			}
		}

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
	}
	return 0
}

func setupGame() {
	// TODO: This will be a menu to start the game
	// But for now we will start the game immediatelly

	// Load the default font
	font_family = rl.LoadFont(fmt.Sprintf("resources/%s", FONT_FAMILY))

	State.SetState(START)
}

// A simple function to restart the game
func reset_game() {
	// todo!()
}

func main() {
	rl.InitWindow(WIDTH, HEIGHT, GAME_TITLE)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	setupGame()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		drawMap()
		render_score()

		// INFO: Debug logger
		_LoggerProcess()

		ret_process := process_state()

		// Exit the game
		if ret_process == 2 {
			break
		}

		rl.EndDrawing()
	}
}
