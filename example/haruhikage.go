package example

import "github.com/Drelf2018/go-beep-player"

var Haruhikage = []beep.Note{
	// 1
	{Key: 2, Tone: 2, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},

	{Key: 2, Tone: 2, Fraction: 8, Dot: 1},
	{Key: 3, Tone: 2, Fraction: 16},
	{Key: 2, Tone: 2, Fraction: 8},
	{Key: 1, Tone: 2, Fraction: 4, Dot: 1},

	{Key: 2, Tone: 2, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},

	{Key: 2, Tone: 2, Fraction: 8, Dot: 1},
	{Key: 3, Tone: 2, Fraction: 16},
	{Key: 2, Tone: 2, Fraction: 8},
	{Key: 1, Tone: 2, Fraction: 4, Dot: 1},

	{Key: 2, Tone: 2, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},

	// 6
	{Key: 2, Tone: 2, Fraction: 8, Dot: 1},
	{Key: 3, Tone: 2, Fraction: 16},
	{Key: 2, Tone: 2, Fraction: 8},
	{Key: 1, Tone: 2, Fraction: 4, Dot: 1},

	{Key: 2, Tone: 2, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 8},

	{Key: 2, Tone: 2, Fraction: 8, Dot: 1},
	{Key: 3, Tone: 2, Fraction: 16},
	{Key: 2, Tone: 2, Fraction: 8},
	{Key: 1, Tone: 2, Fraction: 4},
	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},

	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	// 11
	{Key: 1, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 2, Dot: 1},

	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 8},

	{Key: 6, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},
	{Key: 6, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},

	{Key: 6, Tone: 1, Fraction: 16},
	{Key: 5, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},

	{Key: 3, Tone: 1, Fraction: 4},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 0, Fraction: 8},

	// 16
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},

	{Key: 7, Tone: 0, Fraction: 2, Hold: &beep.Note{Fraction: 8}},
	{Key: 7, Tone: 0, Fraction: 8},

	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 4},
	{Key: 7, Tone: 0, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},

	{Key: 3, Tone: 1, Fraction: 4},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},

	{Key: 7, Tone: 0, Fraction: 2, Hold: &beep.Note{Fraction: 8}},
	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},

	// 21
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	{Key: 1, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 4, Dot: 1, Hold: &beep.Note{Fraction: 4, Dot: 1}},

	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 8},

	{Key: 6, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},
	{Key: 6, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},

	// 26
	{Key: 6, Tone: 1, Fraction: 16},
	{Key: 5, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},

	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 0, Fraction: 8},

	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},

	{Key: 7, Tone: 0, Fraction: 2, Hold: &beep.Note{Fraction: 8}},
	{Key: 7, Tone: 0, Fraction: 8},

	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 4},
	{Key: 7, Tone: 0, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},

	// 31
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},
	{Key: 7, Tone: 0, Fraction: 8},

	{Key: 7, Tone: 0, Fraction: 2, Dot: 1},

	{Key: 5, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},

	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},

	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	// 36
	{Key: 1, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 6, Tone: 0, Fraction: 16},
	{Key: 7, Tone: 0, Fraction: 4, Dot: 1},

	{Key: 5, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},

	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 4},
	{Key: 2, Tone: 1, Fraction: 8},

	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},

	{Key: 1, Tone: 2, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},

	// 41
	{Key: 6, Tone: 1, Fraction: 4},
	{Key: 5, Tone: 1, Fraction: 8},
	{Key: 5, Tone: 1, Fraction: 4, Dot: 1, Hold: &beep.Note{Fraction: 4}},

	{Key: 5, Tone: 1, Fraction: 8},
	{Key: 5, Tone: 1, Fraction: 8},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 16},

	{Key: 3, Tone: 1, Fraction: 4},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 4, Dot: 1},

	{Key: 4, Tone: 0, Fraction: 4, Dot: 1},
	{Key: 5, Tone: 0, Fraction: 4},
	{Key: 6, Tone: 0, Fraction: 8},

	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 4},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 16},

	// 46
	{Key: 5, Tone: 1, Fraction: 4},
	{Key: 5, Tone: 1, Fraction: 16},
	{Key: 6, Tone: 1, Fraction: 16},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 1, Tone: 2, Fraction: 16},
	{Key: 7, Tone: 1, Fraction: 16},

	{Key: 4, Tone: 1, Fraction: 4, Hold: &beep.Note{Fraction: 16}},
	{Key: 6, Tone: 0, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},

	{Key: 2, Tone: 1, Fraction: 4},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 4, Dot: 1},

	{Key: 7, Tone: 0, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 1, Tone: 1, Fraction: 16},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 4},
	{Key: 3, Tone: 1, Fraction: 16},
	{Key: 4, Tone: 1, Fraction: 16},

	// 50
	{Key: 5, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 16, Sharp: 2},
	{Key: 5, Tone: 1, Fraction: 16},
	{Key: 6, Tone: 1, Fraction: 4, Hold: &beep.Note{Fraction: 16}},
	{Key: 2, Tone: 1, Fraction: 16},

	{Key: 2, Tone: 2, Fraction: 8},
	{Key: 2, Tone: 2, Fraction: 8, Dot: 1},
	{Key: 2, Tone: 2, Fraction: 16},
	{Key: 3, Tone: 2, Fraction: 8},
	{Key: 2, Tone: 2, Fraction: 8},
	{Key: 1, Tone: 2, Fraction: 8},

	{Key: 1, Tone: 2, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 16},
	{Key: 6, Tone: 1, Fraction: 16},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 2, Tone: 1, Fraction: 16},
	{Key: 7, Tone: 1, Fraction: 16},

	{Key: 1, Tone: 2, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},

	// 54
	{Key: 1, Tone: 2, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},

	{Key: 1, Tone: 2, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 16},
	{Key: 7, Tone: 1, Fraction: 16},

	{Key: 1, Tone: 2, Fraction: 8, Dot: 1},
	{Key: 2, Tone: 2, Fraction: 16},
	{Key: 1, Tone: 2, Fraction: 8},
	{Key: 7, Tone: 1, Fraction: 4},
	{Key: 7, Tone: 1, Fraction: 8},

	{Key: 6, Tone: 1, Fraction: 4},
	{Key: 5, Tone: 1, Fraction: 8},
	{Key: 5, Tone: 1, Fraction: 4},
	{Key: 4, Tone: 1, Fraction: 8},

	{Key: 4, Tone: 1, Fraction: 4},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	// 59
	{Key: 2, Tone: 1, Fraction: 2, Dot: 1},

	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 3, Tone: 1, Fraction: 8},
	{Key: 2, Tone: 1, Fraction: 8},
	{Key: 1, Tone: 1, Fraction: 8},

	{Key: 7, Tone: 0, Fraction: 2, Dot: 1},
}
