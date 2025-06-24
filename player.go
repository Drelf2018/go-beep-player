package beep

import (
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

//go:linkname modwinmm golang.org/x/sys/windows.modwinmm
var modwinmm *windows.LazyDLL

var (
	procMidiOutOpen     = modwinmm.NewProc("midiOutOpen")
	procMidiOutClose    = modwinmm.NewProc("midiOutClose")
	procMidiOutShortMsg = modwinmm.NewProc("midiOutShortMsg")
)

type Player struct {
	Major      map[int]int
	BPM        int
	Channel    int
	Instrument int
	Velocity   int
	Period     int
	Notes      []Note

	handle    uintptr
	playNotes []playNote
}

type playNote struct {
	begin    uintptr
	end      uintptr
	duration int
	time     time.Duration
}

func (p *Player) Call(a uintptr) error {
	if a == 0 {
		return nil
	}
	_, _, err := procMidiOutShortMsg.Call(p.handle, a)
	return err
}

func (p *Player) Play() {
	ret, _, _ := procMidiOutOpen.Call(uintptr(unsafe.Pointer(&p.handle)), uintptr(0xFFFFFFFF), 0, 0, 0)
	if ret != 0 {
		return
	}
	defer procMidiOutClose.Call(p.handle)

	p.Channel = p.Channel & 0x0F
	p.playNotes = []playNote{{
		end: uintptr(0x80 | p.Channel | (0 << 8) | (p.Velocity << 16)),
	}}

	for _, note := range p.Notes {
		if !note.Natural && note.Sharp == 0 && p.Major != nil {
			note.Sharp = p.Major[note.Key]
		}
		playNote := playNote{
			duration: int(note.Duration() * 60000 / float64(p.BPM*p.Period)),
			time:     time.Duration(note.Duration()*60000/float64(p.BPM)) * time.Millisecond,
		}
		if note.Key != 0 {
			key := note.KeyNum()
			playNote.begin = uintptr(0x90 | p.Channel | (key << 8) | (p.Velocity << 16))
			playNote.end = uintptr(0x80 | p.Channel | (key << 8) | (p.Velocity << 16))
		}
		p.playNotes = append(p.playNotes, playNote)
	}

	windows.TimeBeginPeriod(1)
	defer windows.TimeEndPeriod(1)

	p.Call(uintptr(0xC0 | p.Channel | (p.Instrument << 8)))

	for i := 1; i <= len(p.playNotes); i++ {
		p.Call(p.playNotes[i-1].end)
		if i >= len(p.playNotes) {
			break
		}
		p.Call(p.playNotes[i].begin)
		time.Sleep(p.playNotes[i].time)
	}

	// index := 0
	// usedTick := -1
	// currentTick := 0

	// ticker := time.NewTicker(time.Duration(p.Period) * time.Millisecond)
	// defer ticker.Stop()

	// for range ticker.C {
	// 	if usedTick == currentTick {
	// 		index++
	// 		if index >= len(p.playNotes) {
	// 			break
	// 		}
	// 		go func(i int) {
	// 			// fmt.Printf("index: %v p.playNotes[i-1]: %v p.playNotes[i]: %v\n", i, p.playNotes[i-1], p.playNotes[i])

	// 			p.Call(p.playNotes[i-1].end)
	// 			p.Call(p.playNotes[i].begin)
	// 		}(index)

	// 		// go func(i int) {
	// 		// 	fmt.Printf("index: %v p.playNotes[index]: %v\n", i, p.playNotes[i])
	// 		// 	if i != 0 {
	// 		// 		p.Call(p.playNotes[i-1].end)
	// 		// 	}
	// 		// 	p.Call(p.playNotes[i].begin)
	// 		// }(index)
	// 		// index++
	// 		// if index >= len(p.playNotes) {
	// 		// 	break
	// 		// }
	// 		usedTick = 0
	// 		currentTick = p.playNotes[index].duration
	// 	}
	// 	usedTick++
	// }
}
