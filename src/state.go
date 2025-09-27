package main

import (
	"errors"
)

// Simple state machine to make life more easier
type state struct {
	current_state int
	state_data    map[int]int // TODO: This can be generic for better code's structure
}

func Init(_current_state int) state {
	// p.current_state = _current_state
	return state{
		current_state: _current_state,
	}
}

func (p *state) setState(_State int) {
	p.current_state = _State
}

func (p *state) getState() int {
	return p.current_state
}

func (p *state) setDataFrame(_data_frame map[int]int) {
	p.state_data = _data_frame
}

func (p *state) setData(_user_keys int, _user_data int) error {
	if _, ok := p.state_data[_user_keys]; ok {
		p.state_data[_user_keys] = _user_data
		return nil
	}
	return errors.New("can't find data")
}

func (p *state) getData(key int) int {
	if value, ok := p.state_data[key]; ok {
		return value
	}
	return -1 // TODO: We can use error.new here
}

func (p *state) getDataFrame() map[int]int {
	return p.state_data
}

var State state = Init(IDLE)
