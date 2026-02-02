package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// It's just a simple Tuple, then I will just do the [2]int
type Tuple struct {
	a int
	b int
}

func newTuple(user_a int, user_b int) Tuple {
	return Tuple{
		a: user_a,
		b: user_b,
	}
}

func (p *Tuple) getFirst() int {
	return p.a
}

func (p *Tuple) getSecond() int {
	return p.b
}

func Bool2int(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

/*
		<Summary>
			Notice you need to provide the width & Height of your container for this,
			For example you can use your screen width and height
	  </Summary>
*/
func calc_text_center(text string, temp_width int, temp_height int, fontSize int) (int, int) {
	// Measure the text width
	textWidth := int(rl.MeasureText(text, int32(fontSize)))

	// Calculate position so it’s centered
	var x int = (temp_width / 2) - (textWidth / 2)
	var y int = (temp_height / 2) - (fontSize / 2)

	return x, y
}

// It's just a test function to print logs
func _LoggerProcess() {
	if !DEBUG {
		return
	}
}
