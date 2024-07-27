package example_test

import (
	"testing"

	"github.com/Drelf2018/go-beep-player"
	"github.com/Drelf2018/go-beep-player/example"
)

func TestHaruhikage(t *testing.T) {
	player := beep.Player{BPM: 95, Major: beep.MajorB}
	player.Play(0, 1, 100, example.Haruhikage)
}
