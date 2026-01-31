package main

import (
	"errors"
)

// Simple state machine to make life more easier
type state struct {
	current_state int
	prev_state    int
	state_data    map[int]int // TODO: This can be generic for better code's structure
}

func Init(_current_state int) state {
	// p.current_state = _current_state
	return state{
		current_state: _current_state,
	}
}

func (p *state) SetState(_State int) {
	p.prev_state = p.current_state
	p.current_state = _State
}

func (p *state) PrevState() {
	p.current_state = p.prev_state
	p.prev_state = p.current_state
}

func (p *state) GetState() int {
	return p.current_state
}

func (p *state) GetPrevState() int {
	return p.prev_state
}

func (p *state) SetDataFrame(_data_frame map[int]int) {
	p.state_data = _data_frame
}

func (p *state) SetData(_user_keys int, _user_data int) error {
	if _, ok := p.state_data[_user_keys]; ok {
		p.state_data[_user_keys] = _user_data
		return nil
	}
	return errors.New("can't find data")
}

func (p *state) GetData(key int) int {
	if value, ok := p.state_data[key]; ok {
		return value
	}
	return -1 // TODO: We can use error.new here
}

func (p *state) GetDataFrame() map[int]int {
	return p.state_data
}

var State state = Init(IDLE)
