package beep

import "math"

type Note struct {
	// [1, 7]
	Key int
	// [-4, 4]
	Tone int
	// [-2, 2]
	Sharp   int
	Natural bool
	// {1, 2, 4, 8, 16, ...}
	Fraction int
	// [0, +∞)
	Dot  int
	Hold *Note
}

func (n *Note) KeyNum() int {
	return 58 + 2*n.Key - n.Key/4 + 12*n.Tone + n.Sharp
}

func (n *Note) Duration() float64 {
	if n == nil {
		return 0
	}
	return 4.0/float64(n.Fraction)*2*(1-math.Pow(0.5, 1+float64(n.Dot))) + n.Hold.Duration()
}
