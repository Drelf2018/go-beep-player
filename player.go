package beep

import (
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Player struct {
	Major map[int]int
	BPM   int
}

type playNote struct {
	key      int
	duration time.Duration
}

func (p *Player) Play(channel, instrument, velocity int, notes []Note) {
	var handle uintptr
	winmm := windows.NewLazyDLL("winmm.dll")
	ret, _, _ := winmm.NewProc("midiOutOpen").Call(uintptr(unsafe.Pointer(&handle)), uintptr(0xFFFFFFFF), 0, 0, 0)
	defer winmm.NewProc("midiOutClose").Call(handle)
	if ret != 0 {
		return
	}
	midiOutShortMsg := winmm.NewProc("midiOutShortMsg")

	channel = channel & 0x0F
	playList := make([]playNote, 0, len(notes))
	duration := 60000000 / float64(p.BPM)

	for _, note := range notes {
		if !note.Natural && note.Sharp == 0 && p.Major != nil {
			note.Sharp = p.Major[note.Key]
		}
		playList = append(playList, playNote{
			key:      note.KeyNum(),
			duration: time.Duration(note.Duration()*duration-5250) * time.Microsecond, // why 5250?
		})
	}

	var msg uint32
	msg = uint32(0xC0 | channel | (instrument << 8))
	midiOutShortMsg.Call(handle, uintptr(msg))

	for _, note := range playList {
		msg = uint32(0x90 | channel | (note.key << 8) | (velocity << 16))
		midiOutShortMsg.Call(handle, uintptr(msg))

		time.Sleep(note.duration)

		msg = uint32(0x80 | channel | (note.key << 8) | (velocity << 16))
		midiOutShortMsg.Call(handle, uintptr(msg))
	}
}
