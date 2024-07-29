package beep

import (
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var winmm = windows.NewLazyDLL("winmm.dll")

type Player struct {
	Major      map[int]int
	BPM        int
	Channel    int
	Instrument int
	Velocity   int
	Period     int
	Notes      []Note

	handle          uintptr
	midiOutShortMsg *windows.LazyProc
	playNotes       []playNote
}

type playNote struct {
	begin    uintptr
	end      uintptr
	duration int
}

func (p *Player) Call(a uintptr) {
	p.midiOutShortMsg.Call(p.handle, a)
}

func (p *Player) Play() {
	ret, _, _ := winmm.NewProc("midiOutOpen").Call(uintptr(unsafe.Pointer(&p.handle)), uintptr(0xFFFFFFFF), 0, 0, 0)
	if ret != 0 {
		return
	}

	p.midiOutShortMsg = winmm.NewProc("midiOutShortMsg")

	p.Channel = p.Channel & 0x0F
	p.playNotes = []playNote{{
		end: uintptr(0x80 | p.Channel | (0 << 8) | (p.Velocity << 16)),
	}}

	p.Call(uintptr(0xC0 | p.Channel | (p.Instrument << 8)))

	for _, note := range p.Notes {
		if !note.Natural && note.Sharp == 0 && p.Major != nil {
			note.Sharp = p.Major[note.Key]
		}
		key := note.KeyNum()
		p.playNotes = append(p.playNotes, playNote{
			begin:    uintptr(0x90 | p.Channel | (key << 8) | (p.Velocity << 16)),
			end:      uintptr(0x80 | p.Channel | (key << 8) | (p.Velocity << 16)),
			duration: int(note.Duration() * 60000 / float64(p.BPM*p.Period)),
		})
	}

	defer winmm.NewProc("midiOutClose").Call(p.handle)

	windows.TimeBeginPeriod(1)
	defer windows.TimeEndPeriod(1)

	ticker := time.NewTicker(time.Duration(p.Period) * time.Millisecond)
	defer ticker.Stop()

	index := 0
	usedTick := 0
	currentTick := 0
	for range ticker.C {
		if usedTick == currentTick {
			index++
			if index >= len(p.playNotes) {
				break
			}
			go func() {
				p.Call(p.playNotes[index-1].end)
				p.Call(p.playNotes[index].begin)
			}()
			usedTick = 0
			currentTick = p.playNotes[index].duration
		}
		usedTick++
	}
}
