package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Manage key bindings

// Basic movements enums
const (
	RIGHT = iota
	LEFT
	UP
	BOTTOM
	PLAY
)

// Default key bindings
var keys map[int]int = map[int]int{
	RIGHT:  rl.KeyRight,
	LEFT:   rl.KeyLeft,
	UP:     rl.KeyDown,
	BOTTOM: rl.KeyUp,
	PLAY:   rl.KeySpace,
}

func Get_key(key int) int {
	if value, ok := keys[key]; ok {
		return value
	}
	return -1
}
