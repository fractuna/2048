package main

// It's just a simple Tuple, then I will just do the int
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
