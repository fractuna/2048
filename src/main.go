package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	score         int  = 0
	l_item        int  = 8 // based on the default game's map
	items_checked bool = false
	font_family   rl.Font
	isClean       bool = false
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

// var tileMap [MAX_X][MAX_Y]int = [MAX_X][MAX_Y]int{
// 	{2, 4, 2, 4},
// 	{4, 2, 4, 2},
// 	{2, 4, 2, 4},
// 	{4, 2, 4, 2},
// }

var tileMap [MAX_X][MAX_Y]int = [MAX_X][MAX_Y]int{
	{2, 0, 4, 4},
	{0, 8, 0, 4},
	{4, 2, 0, 8},
	{0, 0, 16, 16},
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
	fmt.Println("GAME OVER!")
	x, y := calc_text_center(GAME_OVER_TEXT, rl.GetScreenWidth(), rl.GetScreenHeight(), H1)
	rl.DrawTextEx(font_family, GAME_OVER_TEXT, rl.NewVector2(float32(x), float32(y)), H1, 4, rl.Black)
	// rl.DrawText(GAME_OVER_TEXT, (int32(rl.GetScreenWidth()/2))-(rl.MeasureText(GAME_OVER_TEXT, DEFAULT_FONT_SIZE)), 510/2, DEFAULT_FONT_SIZE, rl.Black)
}

func process_state() int {
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
		fmt.Println("MOVE LEFT")
		isClean = Move_v(3, -1, 0)
		fmt.Println(isClean)
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
		isCleanZero := Move_zero_v(0, +1, MAX_X-1)
		if isCleanZero && isClean {
			State.SetState(JUST_FINISH)
		} else {
			State.PrevState()
		}

	case MOVE_ZERO_LEFT:
		isCleanZero := Move_zero_v(3, -1, 0)
		fmt.Println(isClean, isCleanZero)
		if isCleanZero && isClean {
			State.SetState(JUST_FINISH)
		} else {
			State.PrevState()
		}

	case MOVE_ZERO_DOWN:
		isCleanZero := Move_zero_h(3, -1, 0)
		// isCleanZero := Move_zero_v(0, +1, MAX_X-1)
		if isCleanZero && isClean {
			State.SetState(JUST_FINISH)
		} else {
			State.PrevState()
		}

	case MOVE_ZERO_UP:
		isCleanZero := Move_zero_h(0, +1, MAX_Y-1)
		if isCleanZero && isClean {
			State.SetState(JUST_FINISH)
		} else {
			State.PrevState()
		}
	case JUST_FINISH:
		if l_item >= L_MAX {
			if items_checked == false {
				// Check if the player can't do any movement
				State.SetState(CEHCK_ITEMS)
				if possible_move() == false {
					fmt.Println("Player can't move anymore")
					State.SetState(LOSE)
					// drawLosePopup()
				} else {
					fmt.Println("Player can still move")
					State.SetState(IDLE)
				}
				items_checked = true
			}
		} else {
			// TODO: Don't add new item when there is 2 in a row clean movement
			Add_item()
			State.SetState(IDLE)
		}
	case LOSE:
		fmt.Println("LOSE STATE")
		game_over()
		// return 2
	case IDLE:

		isClean = false
		// if l_item >= L_MAX {
		// 	if items_checked == false {
		// 		// Check if the player can't do any movement
		// 		State.SetState(CEHCK_ITEMS)
		// 		if possible_move() {
		// 			State.SetState(LOSE)
		// 			// drawLosePopup()
		// 		} else {
		// 			State.SetState(IDLE)
		// 		}
		// 		items_checked = true
		// 	}
		// }

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

func drawLosePopup() {
	// Popup dimensions
	popupWidth := int32(300)
	popupHeight := int32(150)
	screenWidth := rl.GetScreenWidth()
	screenHeight := rl.GetScreenHeight()
	x := (int32(screenWidth) - popupWidth) / 2
	y := (int32(screenHeight) - popupHeight) / 2

	// Draw semi-transparent background
	rl.DrawRectangle(x, y, popupWidth, popupHeight, rl.Fade(rl.RayWhite, 0.95))

	// Draw border
	rl.DrawRectangleLines(x, y, popupWidth, popupHeight, rl.Black)

	// Draw lose text (split into two lines for proper centering)
	fontSize := int32(24)
	loseText1 := "Game Over!"
	loseText2 := "Press R to Restart"
	textWidth1 := rl.MeasureText(loseText1, fontSize)
	textWidth2 := rl.MeasureText(loseText2, fontSize)
	rl.DrawText(loseText1, x+(popupWidth-textWidth1)/2, y+40, fontSize, rl.Red)
	rl.DrawText(loseText2, x+(popupWidth-textWidth2)/2, y+80, fontSize, rl.Red)
}

func setupGame() bool {
	// TODO: This will be a menu to start the game
	// But for now we will start the game immediatelly

	// Load the default font
	// font_family = rl.LoadFont(fmt.Sprintf("resources/%s", FONT_FAMILY))

	font_asset, err := Asset(fmt.Sprintf("resources/%s", FONT_FAMILY))
	if err != nil {
		fmt.Println("ERROR: Can't load the font, maybe rollback to default font")
		return false
	}

	font_family = rl.LoadFontFromMemory(".ttf", font_asset, 256, []rune("abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTYVWXYZz0123456789.!: "))

	State.SetState(START)
	return true
}

// A simple function to restart the game
func reset_game() {
	// todo!()
}

func main() {
	rl.InitWindow(WIDTH, HEIGHT, GAME_TITLE)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	if !setupGame() {
		os.Exit(1)
	}

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

	rl.UnloadFont(font_family)
}
