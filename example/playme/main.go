package main

import (
	"github.com/Drelf2018/go-beep-player"
	"github.com/Drelf2018/go-beep-player/example"
)

func main() {
	player := &beep.Player{
		Major:      beep.MajorB,
		BPM:        95,
		Channel:    0,
		Instrument: 0,
		Velocity:   127,
		Period:     3,
		Notes:      example.Haruhikage,
	}
	player.Play()
}
